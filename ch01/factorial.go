package main

import "fmt"

func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n-1)
}

func facItr(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	fmt.Println(fac(5))
	fmt.Println(facItr(10))
}
