// Uniq removes adjacent duplicates from standard input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var previous, word string
	for scanner.Scan() {
		word = scanner.Text()
		if word != previous {
			previous = word
			fmt.Println(word)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "got error: %v\n", err)
		os.Exit(1)
	}
}
