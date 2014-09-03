package main

import "fmt"

// START OMIT
const (
	_ = 1 << (iota * 10)
	KB
	MB
	GB
)

// END OMIT

func main() {
	fmt.Printf("KB =  %v\n", KB)
	fmt.Printf("MB =  %v\n", MB)
	fmt.Printf("GB =  %v\n", GB)

}
