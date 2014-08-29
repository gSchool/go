package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var (
		name string
		now  = time.Now()
	)

	name = "Billy Bob"
	kitchenTime := now.Format(time.Kitchen)
	// END OMIT
	fmt.Printf("name = %q with a type of %T\n", name, name)
	fmt.Printf("kitchenTime is %v", kitchenTime)
}
