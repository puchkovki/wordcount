// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		words := strings.Fields(args[0])
		fmt.Println(len(words))
	} else {
		fmt.Println(0)
	}
}
