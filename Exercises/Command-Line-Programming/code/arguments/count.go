package main

import (
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args[1:])
	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)
}
