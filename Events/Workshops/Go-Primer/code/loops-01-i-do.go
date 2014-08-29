package main

import "fmt"

func main() {
	// START OMIT
	for i := 0; i <= 5; i++ {
		fmt.Printf("Is %d odd? %v\n", i, odd(i))
	}

	for i := 5; i >= 0; i-- {
		fmt.Printf("Is %d odd? %v\n", i, odd(i))
	}
	// END OMIT
}

func odd(i int) bool {
	switch i {
	case 0:
		// Special case, 0 is actual considered "even"
		return false
	default:
		return i%2 == 1
	}
}
