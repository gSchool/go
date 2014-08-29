package main

import "fmt"
import "os"

func main() {
	// START OMIT
	if f, err := openFile(""); err != nil {
		fmt.Printf("Error opening file: %q\n", err)
	} else {
		fmt.Printf("Opened file %v\n", f)
	}
	// END OMIT
}

// Mock open file for demonstration purposes
func openFile(name string) (os.File, error) {
	if len(name) == 0 {
		return os.File{}, fmt.Errorf("File name can not be blank")
	} else {
		return os.File{}, nil
	}
}
