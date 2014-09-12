package main

import (
	"fmt"
	"os"
)

// fake out the args for purposes of the present tool
func init() {
	os.Args = append(os.Args, "-local", "u=admin", "--help")
}

//START OMIT
func main() {
	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
}

//END OMIT
