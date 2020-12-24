package main

import "fmt"

// selection sort
func sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}
}

func main() {
	list := []int{9, 3, 2, -5, 13, -26, 291}
	fmt.Println("original:", list)

	sort(list)
	fmt.Println("sorted:", list)
}
