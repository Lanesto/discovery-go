package main

import "fmt"

func sing(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%d %d\n", i*10, (i+1)*10)
	}
}

func main() {
	sing(5)
}
