// Dup3 prints the text of each line that appears more than
// once in a list of named files, preceded by its count.
// This version loads the entire file into memory, rather than
// streaming it (as dup1 and dup2 do)
// This is similar to the UNIX program uniq.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		fmt.Printf("dup3: %s\n", filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			fmt.Printf("dup3: line=%q\n", line)
			if line == "" {
				continue
			}
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
