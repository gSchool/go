package main

import "fmt"

// START OMIT
const (
	_   int = iota // Skip the first value of 0
	Foo            // Foo = 1
	Bar            // Bar = 2
	Bin            // Bin = 3
	_
	_
	Baz // Baz = 6
)

// END OMIT

func main() {
	fmt.Printf("The value of Foo is %v\n", Foo)
	fmt.Printf("The value of Bar is %v\n", Bar)
	fmt.Printf("The value of Bin is %v\n", Bin)
	fmt.Printf("The value of Baz is %v\n", Baz)
}
