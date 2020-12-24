package main

import "fmt"

func fibo(n int) int {
	// fibo(0), fibo(1)
	p, q := 0, 1
	for i := 0; i < n; i++ {
		// fibo(n-2), fibo(n-1)
		p, q = q, p+q
	}
	return p
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(10))
}
