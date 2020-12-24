package main

import (
	"errors"
	"fmt"
)

func push(q *[]int, value int) {
	*q = append(*q, value)
}

func pop(q *[]int) int {
	if len(*q) == 0 {
		panic(errors.New("queue is empty"))
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func main() {
	queue := &[]int{}
	archive := queue

	fmt.Println(queue)
	push(queue, 1)

	fmt.Println(queue)
	push(queue, 2)

	fmt.Println(queue)
	_ = pop(queue)

	fmt.Println(queue)
	push(queue, 3)

	fmt.Println(queue)
	_ = pop(queue)

	fmt.Println(queue)
	fmt.Println(archive)
}
