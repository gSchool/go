package main

import "fmt"

func main() {
	// START OMIT
	tag := 3

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
