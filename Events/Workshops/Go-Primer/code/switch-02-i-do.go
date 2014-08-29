package main

import "fmt"

func main() {
	// START OMIT
	x, y := 3, 4

	switch {
	case x < y:
		fmt.Printf("%d is less than %d\n", x, y)
	case x > y:
		fmt.Printf("%d is more than %d\n", x, y)
	default:
		fmt.Printf("%d is equal to %d\n", x, y)
	}
	// END OMIT

}
