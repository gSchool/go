package main

import "fmt"

func main() {
	//START OMIT
	var guessColor string
	const favColor string = "blue"
	for {
		println("Guess my favorite color:")
		if _, err := fmt.Scanf("%s", &guessColor); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		if favColor == guessColor {
			fmt.Printf("%q is my favorite color!", favColor)
			return
		}
		fmt.Printf("Sorry, %q is not my favorite color.  Guess again.\n", guessColor)
	}
	// END OMIT
}
