package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var (
		name string              // basic declaration with zero value of ""
		x    int    = 5          // declaration and initialization
		now         = time.Now() // initialization with inferred type
	)

	// change the value of the variable
	name = "Billy Bob"

	// declare & initialize new variable and infer type
	kitchenTime := now.Format(time.Kitchen)

	// END OMIT
	fmt.Printf("name = %q with a type of %T\n", name, name)
	fmt.Printf("x = %d with a type of %T\n", x, x)
	fmt.Printf("kitchenTime is %v", kitchenTime)
}
