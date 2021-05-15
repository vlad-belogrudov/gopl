// charcount produces statistics for a given UTF-8 text
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	letters := 0
	digits := 0
	punctuations := 0
	spaces := 0
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			letters++
		} else if unicode.IsDigit(r) {
			digits++
		} else if unicode.IsPunct(r) {
			punctuations++
		} else if unicode.IsSpace(r) {
			spaces++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i := 1; i < len(utflen); i++ {
		fmt.Printf("%d\t%d\n", i, utflen[i])
	}
	fmt.Printf("letters: %d\n", letters)
	fmt.Printf("digits: %d\n", digits)
	fmt.Printf("punctuations: %d\n", punctuations)
	fmt.Printf("spaces: %d\n", spaces)
	fmt.Printf("invalid: %d\n", invalid)
}
