package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args[0] // executing program name
	argCount := len(os.Args[1:])
	fmt.Printf("Program Name: %s\n", cmd)
	fmt.Printf("Total Arguments: %d\n", argCount)
	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
}
