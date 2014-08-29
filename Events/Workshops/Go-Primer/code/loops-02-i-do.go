package main

import "fmt"

func main() {
	a := 1
	b := 10

	// START OMIT
	for a < b {
		a += 1
		fmt.Println(a)
	}
	// END OMIT

}
