package main

import "fmt"

// START OMIT
const (
	read   = 1 << iota // 00000001 = 1
	write              // 00000010 = 2
	remove             // 00000100 = 4

)

// END OMIT

func main() {
	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)

}
