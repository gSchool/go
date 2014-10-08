package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//START OMIT
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("	cat file1 file2 ...\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, fn := range flag.Args() {
		if f, e := os.Open(fn); e != nil {
			panic(e)
		} else {
			_, err := io.Copy(os.Stdout, f)
			if err != nil {
				panic(err)
			}
		}
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
	}
	//END OMIT

}
