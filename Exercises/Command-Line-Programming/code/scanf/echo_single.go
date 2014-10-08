package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// START OMIT
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
		return
	}
	// END OMIT
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
