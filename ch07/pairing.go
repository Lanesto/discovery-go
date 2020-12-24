package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Value   int
	ReplyTo chan Response
}

type Response struct {
	Value    int
	WorkerID int
}

func CreateWorker(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.ReplyTo)
			req.ReplyTo <- Response{req.Value * 2, workerID}
		}(req)
	}
}

func main() {
	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go CreateWorker(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			postbox := make(chan Response)
			reqs <- Request{i, postbox}
			fmt.Println(i, "=>", <-postbox)
		}(i)
	}
	wg.Wait()
}
