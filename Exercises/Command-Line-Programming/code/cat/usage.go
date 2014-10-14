package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("	cat file1 file2 ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}
