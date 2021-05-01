// Uniq removes adjacent duplicates from []string
package main

import (
	"fmt"
	"os"
)

func main() {
	words := os.Args[1:]
	i := 0
	var previous string
	for _, word := range words {
		if word != previous {
			words[i] = word
			previous = word
			i++
		}
	}
	fmt.Println(words[:i])
}
