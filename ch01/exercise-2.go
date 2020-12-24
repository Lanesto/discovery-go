package main

import "fmt"

func solve(n, from, to, buffer int) {
	if n == 1 {
		fmt.Printf("%d -> %d\n", from, to)
		return
	}
	solve(n-1, from, buffer, to)
	fmt.Printf("%d -> %d\n", from, to)
	solve(n-1, buffer, to, from)
}

func hanoi(n int) {
	solve(n, 1, 2, 3)
}

func main() {
	n := 3
	fmt.Println("Number of disks:", n)
	hanoi(n)
}
