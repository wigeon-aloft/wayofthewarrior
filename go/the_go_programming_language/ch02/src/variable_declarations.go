package main

import (
	"os"
)

func main() {
	/* A full variable declaration looks like the following:
		var name type = expression

	`= expression` or `type` can be omitted, but not both.
	*/

	// Variable declaration
	var s1 string = "hello, world!"
	var s2 string

	// Multiple variable declaration
	var i, j, k, l int              // Same type: int
	var b, f, s = true, 2.3, "four" // different types: bool, float64, string

	// Setting variables from a function that returns multiple values
	var f1, err1 = os.Open(name)

	// Short variable declaration
	// Used to declare local variables within functions
	// The type is determined by the expression
	t := 0.0

	// Multiple variable declaration
	m, n := 0, 1

	// Setting variables from function that returns multiple values
	f2, err2 := os.Open(name)

	/* The short variable declaration := can also be used as an assignment operator,
	if some of the variables on the left hand side of the operator have already been assigned.

	Note that at least one new value needs to be declared, in this case that's `f3`.
	*/
	f3, err := os.Open(name)
}
