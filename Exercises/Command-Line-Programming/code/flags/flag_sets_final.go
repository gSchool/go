package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	os.Args = append(os.Args, []string{"--help"}...)
}

// START OMIT
func main() {
	var (
		cmd  string = "website"
		port int    = 8000
		help bool
		fs   = flag.NewFlagSet("default", flag.ExitOnError)
	)

	fs.BoolVar(&help, "help", false, "prints out usage")
	fs.StringVar(&cmd, "cmd", cmd, "the command to run")
	fs.IntVar(&port, "port", port, "the port to run on")
	fs.Parse(os.Args[1:])

	if help {
		fmt.Println("Usage:")
		fs.PrintDefaults()
		return
	}
	if cmd != "website" {
		fmt.Printf("Running %s", cmd)
		// Do something like run a website
		return
	}
}

// END OMIT
