package seq

// Fib returns nth (n >= 0) fibonacci number
func Fib(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p + q
	}
	return p
}
