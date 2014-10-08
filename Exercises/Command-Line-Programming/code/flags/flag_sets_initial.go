package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	os.Args = append(os.Args, []string{"-h"}...)
}

func main() {
	var (
		cmd  string = "website"
		port int    = 8000
	)

	// START OMIT
	fs := flag.NewFlagSet("default", flag.ExitOnError)

	fs.StringVar(&cmd, "cmd", cmd, "the command to run")
	fs.IntVar(&port, "p", port, "the port to run on")

	// END OMIT
	fs.Parse(os.Args[1:])

	fmt.Printf("Running %q on port %d", cmd, port)
}
