// This module demonstrates a basic usage of goroutines (with the `go` keyword)
// The script will display a spinner in the console whilst the slow fibonacci
// number generator runs in the background.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Start the spinner in a goroutine using the `go` keyword
	// Goroutines are similar to threads in other programming languages
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	// All goroutines are exited when the main goroutine exits
	// so the spinner stops once the fibonacci number is output
	// and we get here.
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
