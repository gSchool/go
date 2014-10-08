package main

import "fmt"

func main() {
	//START OMIT
	var color string
	for {
		println("Guess my favorite color:")
		if _, err := fmt.Scanf("%s", &color); err != nil {
			println(err)
			return
		}
		if color == "blue" {
			println("Blue is my favorite color!")
			return
		}
		fmt.Printf("Sorry, %s is not my favorite color.  Guess again.\n", color)
	}
	// END OMIT
}
