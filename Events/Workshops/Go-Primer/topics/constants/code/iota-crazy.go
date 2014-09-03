package main

import "fmt"

// START OMIT
const (
	tomato, apple int = iota + 1, iota + 2
	orange, chevy
	ford, _
)

// END OMIT

func main() {
	fmt.Printf("The value of tomato is %v\n", tomato)
	fmt.Printf("The value of apple is %v\n", apple)
	fmt.Printf("The value of orange is %v\n", orange)
	fmt.Printf("The value of chevy is %v\n", chevy)
	fmt.Printf("The value of ford is %v\n", ford)
}
