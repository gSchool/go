package main

import "fmt"

func main() {
	// START OMIT
	var char rune = 'A' // int32 value of 65

	letter := string(char)
	i := int(char)

	name := "Bill"
	bytes := []byte(name)
	// END OMIT

	fmt.Printf("char has a value of %[1]v and a type of %[1]T\n", char)
	fmt.Printf("letter has a value of %[1]v and a type of %[1]T\n", letter)
	fmt.Printf("i has a value of %[1]v and a type of %[1]T\n", i)
	fmt.Println()
	fmt.Printf("name has a value of %[1]v and a type of %[1]T\n", name)
	fmt.Printf("bytes  a value of %[1]v and a type of %[1]T\n", bytes)
}
