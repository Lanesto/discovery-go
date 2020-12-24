package eval

import (
	"fmt"
	"strings"
)

var opDef = OpDef{
	"**": func(a, b int) int {
		if a == 1 {
			return 1
		}
		if b < 0 {
			return 0
		}
		ret := 1
		for i := 0; i < b; i++ {
			ret *= a
		}
		return ret
	},
	"*":   func(a, b int) int { return a * b },
	"/":   func(a, b int) int { return a / b },
	"mod": func(a, b int) int { return a % b },
	"+":   func(a, b int) int { return a + b },
	"-":   func(a, b int) int { return a - b },
}

var precDef = PriorityDef{
	"**":  NewSet(),
	"*":   NewSet("**", "*", "/", "mod"),
	"/":   NewSet("**", "*", "/", "mod"),
	"mod": NewSet("**", "*", "/", "mod"),
	"+":   NewSet("**", "*", "/", "mod", "+", "-"),
	"-":   NewSet("**", "*", "/", "mod", "+", "-"),
}

func ExampleEvaluator() {
	e := NewEvaluator(opDef, precDef)
	run := func(expr string) int {
		n, _ := e.Eval(expr)
		return n
	}
	fmt.Println(run("5 "))
	fmt.Println(run("1  + 2"))
	fmt.Println(run("1 - 2 - 4  "))
	fmt.Println(run("( 3 - 2 **   3 ) * ( -2 )"))
	fmt.Println(run("3 * ( 3 +  1 * 3 ) / ( -2 )"))
	fmt.Println(run("3 * ( (  3 + 1 ) * 3 ) / 2"))
	fmt.Println(run("1 + 2 ** 10 * 2"))
	fmt.Println(run("2  ** 3 mod 3"))
	fmt.Println(run("2  ** 2 **  3"))

	phrase := strings.Join([]string{
		"다들 그 동안 고생이 많았다.",
		"첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
		"둘째는 임야 { 10 ** 5 mod 7777 }평을 가져라.",
		"막내는 { 10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고",
		"배기량 { 711 * 8 / 9}cc의 경운기를 갖거라.",
	}, "\n")
	fmt.Println(e.Parse(phrase))
	// Output:
	// 5
	// 3
	// -5
	// 10
	// -9
	// 18
	// 2049
	// 2
	// 256
	// 다들 그 동안 고생이 많았다.
	// 첫째는 분당에 있는 48평 아파트를 갖거라.
	// 둘째는 임야 6676평을 가져라.
	// 막내는 3324평 임야를 갖고
	// 배기량 632cc의 경운기를 갖거라.
}
