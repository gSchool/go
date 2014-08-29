package main

import "fmt"

func main() {
	compare(1, 2)
	compare(2, 1)
	compare(1, 1)

}

func compare(x, y int) {
	// START OMIT
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
