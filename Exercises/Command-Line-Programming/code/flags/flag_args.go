package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	os.Args = append(os.Args, "-p=9000", "foo=10", "bar")

}

// START OMIT
func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000.")
	flag.Parse()

	fmt.Printf("port = %d\n", port)
	fmt.Printf("other args: %+v\n", flag.Args())
}

// END OMIT
