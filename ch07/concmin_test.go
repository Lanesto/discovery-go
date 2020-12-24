package concmin

import (
	"fmt"
	"testing"
)

var sample = []int{
	83, 46, 49, 23, 92,
	48, 39, 91, 44, 99,
	25, 42, 74, 56, 23,
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(sample)
	}
}

func BenchmarkParallelMin_3GR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelMin(sample, 3)
	}
}

func BenchmarkParallelMin_5GR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelMin(sample, 5)
	}
}

func BenchmarkParallelMin_10GR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelMin(sample, 10)
	}
}

func BenchmarkParallelMin_15GR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelMin(sample, 15)
	}
}

func ExampleMin() {
	fmt.Println(Min(sample))
	// Output: 23
}

func ExampleParallelMin() {
	fmt.Println(ParallelMin(sample, 5))
	// Output: 23
}
