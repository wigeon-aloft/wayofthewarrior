// Echo1 prints its command-line arguments
// This is a simple, but relatively inefficient implementation of the UNIX echo command
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// Go's only loop is the for loop
	// Below is an instance of a for loop with the initialisation, condition, and post statments
	for i := 1; i < len(os.Args); i++ {
		// Note that the '+' operator in Go concats strings
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
