// Echo3 prints its command-line arguments
// This version uses the `Join` function from the `strings` package
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
