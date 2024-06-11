// Ftoc prints two Fahrenheit-to-Celcius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

// Moved Fahrenheit-to-Celcius conversion to a function for reusability
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
