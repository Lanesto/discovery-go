package main

import (
	"fmt"
	"sort"
)

// binary search
func indexOf(sorted []string, str string) int {
	lo, hi := 0, len(sorted)
	for {
		mid := int((hi + lo) / 2)
		if str < sorted[mid] {
			hi = mid
		} else if str == sorted[mid] {
			return mid
		} else {
			lo = mid
		}
		if mid == int((hi+lo)/2) {
			break
		}
	}
	return -1
}

func main() {
	slice := []string{"go", "cpp", "python", "rust", "java", "node", "haskell"}
	sort.Strings(slice)
	fmt.Println("strings:", slice)

	lookups := []string{"python", "go", "node", "swift"}
	for _, str := range lookups {
		idx := indexOf(slice, str)
		fmt.Printf("index of %s: %d\n", str, idx)
	}
}
