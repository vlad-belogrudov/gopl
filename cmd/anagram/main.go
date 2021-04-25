// Anagram checks if two strings are reverse of each other
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments, must be 2 strings")
		os.Exit(1)
	}
	a := []rune(os.Args[1])
	b := []rune(os.Args[2])
	if len(a) != len(b) {
		fmt.Println("Strings have different length")
		return
	}
	for i, end := 0, len(a)-1; i <= end; i++ {
		if a[i] != b[end-i] {
			fmt.Println("Strings differ")
			return
		}
	}
	fmt.Println("Strings are reverse-equal")
}
