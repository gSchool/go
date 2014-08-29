package main

import "fmt"

// START OMIT
const (
	Apple int = iota
	Orange
	Banana
)

// END OMIT

func main() {
	fmt.Printf("The value of Apple is %v\n", Apple)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Banana is %v\n", Banana)
}
