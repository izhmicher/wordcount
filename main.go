// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1]
	wordcount := len(strings.Split(words, " "))
	fmt.Println(wordcount)
}
