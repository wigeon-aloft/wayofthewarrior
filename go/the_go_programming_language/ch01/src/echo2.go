// Echo2 prints its command-line arguments
// This version of echo uses the ':=' operator to define variables and
// a for loop that iterates over a 'range' of values.
package main

import (
	"fmt"
	"os"
)

func main() {
	// Note that the := initialisation operator can only be used in functions.
	// var must be used at the package level.
	s, sep := "", ""

	// When iterating over a range, Go returns the index and the element.
	// In this instance, we don't need the index, so we assign it to the
	// blank identifier '_' (underscore).
	for _, arg := range os.Args[1:] {
		// Note that the '+' operator in Go concats strings
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
