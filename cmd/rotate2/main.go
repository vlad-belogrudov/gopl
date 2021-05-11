// TODO: Rotate string left or right a number of positions
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: rotate <string> <shift>")
		os.Exit(1)
	}
	shift, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Shift must be an integer")
		os.Exit(1)
	}
	line := []rune(os.Args[1])
	Rotate(line, shift)
	fmt.Println(string(line))
}

// Rotate accepts slice of runes and shifts elements
func Rotate(line []rune, shift int) {
	length := len(line)
	shift %= length
	if shift < 0 {
		shift = length + shift
	}

	for start, finish := 0, length; length > 1; length = finish - start {
		// exchange minimum number of chars, e.g. len = 5, rotate = 3 =>
		// change 2 from left and right
		change := length - shift
		if change > shift {
			change = shift
		}
		if change == 0 {
			return
		}
		// exchange runes [start:start+change] and [finish - change; finish]
		for i, j := start, finish-change; i < start+change; i, j = i+1, j+1 {
			line[i], line[j] = line[j], line[i]
		}
		// calculate left or right part to rotate further
		if change < shift {
			start += change
			shift = length - 2*change
		} else {
			finish -= change
			shift = change
		}
	}
}
