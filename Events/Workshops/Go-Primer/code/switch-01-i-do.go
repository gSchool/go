package main

import "fmt"

func main() {
	matches(0)
	matches(1)
	matches(3)
}

func matches(tag int) {
	// START OMIT
	switch tag {
	default:
		fmt.Println("Couldn't match anything")
	case 1:
		fmt.Println("I'm one!")
	case 3:
		fmt.Println("I'm three!")
	}
	// END OMIT
}
