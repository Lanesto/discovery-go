package main

import "fmt"

var (
	start = rune(44032)
	end   = rune(55204)
)

const numConsonants = 28

func hasConsonantSuffix(s string) bool {
	var last rune
	for _, r := range s {
		last = r
	}
	if start <= last && last < end {
		idx := int(last - start)
		return (idx % numConsonants) != 0
	}
	return false
}

func main() {
	fruits := [...]string{"사과", "바나나", "토마토", "수박", "파인애플"}
	for _, fruit := range fruits {
		var conn string
		if hasConsonantSuffix(fruit) {
			conn = "은"
		} else {
			conn = "는"
		}
		fmt.Printf("%s%s 맛있다.\n", fruit, conn)
	}
}
