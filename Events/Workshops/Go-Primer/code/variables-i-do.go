package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var (
		name              string              // basic declaration with zero value of ""
		i, j              int                 // multiple declarations of same type
		x                 int    = 5          // declaration and initialization
		freezing, boiling int    = 32, 212    // multiple declarations and initialization
		now                      = time.Now() // initialization with inferred type
	)

	// change the value of the variable
	name = "Billy Bob"

	// declare & initialize new variable and infer type
	kitchenTime := now.Format(time.Kitchen)

	// END OMIT
	fmt.Printf("name = %q with a type of %T\n", name, name)
	fmt.Printf("x = %d with a type of %T\n", x, x)
	fmt.Println()
	fmt.Printf("now is %[1]v with a type of %[1]T\n", now)
	fmt.Printf("kitchenTime is %[1]v with a type of %[1]T\n", kitchenTime)
	fmt.Println()
	fmt.Printf("i, j = %d, %d\n", i, j)
	fmt.Printf("freezing, boiling = %d, %d\n", freezing, boiling)
}
