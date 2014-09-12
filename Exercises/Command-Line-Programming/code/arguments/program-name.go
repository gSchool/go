package main

import (
	"fmt"
	"os"
)

func main() {
	// Program Name is the first argument
	cmd := os.Args[0]

	fmt.Printf("Program Name: %s\n", cmd)
}
