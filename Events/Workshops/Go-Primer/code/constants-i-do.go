package main

import "fmt"

// START OMIT
const (
	foo = "A"
	bar = 'A' // any guesses what this will be?
	bin = 2
)

// END OMIT

func main() {
	fmt.Printf("foo is a %T with the value of  %q\n", foo, foo)
	fmt.Printf("bar is a %T with the value of  %v\n", bar, bar)
	fmt.Printf("bin is a %T with the value of  %v\n", bin, bin)

	// And for fun
	fmt.Printf("%T %v\n", bar, string(bar))

}
