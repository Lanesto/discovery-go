package main

import (
	"fmt"
	"sync"
)

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range ins {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for n := range in {
				out <- n
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

func main() {
	cut := func(in <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for n := range in {
				fmt.Println("cut", n)
				c <- n
			}
		}()
		return c
	}
	draw := func(in <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for n := range in {
				fmt.Println("draw", n)
				c <- n
			}
		}()
		return c
	}
	paint := func(in <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for n := range in {
				fmt.Println("paint", n)
				c <- n
			}
		}()
		return c
	}
	decorate := func(in <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for n := range in {
				fmt.Println("decorate", n)
				c <- n
			}
		}()
		return c
	}
	box := func(in <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for n := range in {
				fmt.Println("box", n)
				c <- n
			}
		}()
		return c
	}

	in := make(chan int)
	/*
		              ↗ draw → paint → decorate ↘
			  in → cut →         ...              → box → out
		              ↘         ...             ↗
	*/
	out := Chain(cut, Distribute(Chain(draw, paint, decorate), 10), box)(in)
	go func() {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}()

	for n := range out {
		fmt.Println("main", n)
	}

	fmt.Println("end")
	// Not reachable
	// close(in)
}
