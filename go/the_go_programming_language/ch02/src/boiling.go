/*
Boiling prints the boiling point of water
The purpose of this file is to demostrate declarations
*/
package main

import "fmt"

// package level declaration
const boilingF = 212.0

func main() {
	// f and c are local to the function main()
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
