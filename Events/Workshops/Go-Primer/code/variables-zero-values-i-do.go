package main

import "fmt"

// START OMIT
var (
	name  string
	x     int
	ready bool
	price float32
	char  rune
)

// END OMIT

func main() {
	fmt.Printf("name has a value of %[1]q and a type of %[1]T\n", name)
	fmt.Printf("x has a value of %[1]d and a type of %[1]T\n", x)
	fmt.Printf("ready has a value of %[1]v and a type of %[1]T\n", ready)
	fmt.Printf("price has a value of %[1]v and a type of %[1]T\n", price)
	fmt.Printf("char has a value of %[1]v and a type of %[1]T\n", char)
}
