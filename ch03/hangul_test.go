package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	fmt.Println(HasConsonantSuffix("아무개"))
	fmt.Println(HasConsonantSuffix("각다귀"))
	fmt.Println(HasConsonantSuffix("그램"))
	// Output:
	// false
	// true
	// false
	// false
	// false
	// true
}
