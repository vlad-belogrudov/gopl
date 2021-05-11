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
	shift %= len(line)
	if shift < 0 {
		shift = len(line) - shift
	}
	Rotate(line, shift)
	fmt.Println(string(line))
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Rotate(line []rune, shift int) {
	length := len(line)

	change := MinInt(shift, length-shift)
	fmt.Println(string(line), shift, change)
	if length == 0 || change == 0 {
		return
	}
	for i, j := 0, length-change; i < change; i, j = i+1, j+1 {
		line[i], line[j] = line[j], line[i]
	}
	if change < shift {
		fmt.Println("right part")
		Rotate(line[change:], length-2*change)
	} else {
		fmt.Println("left part")
		Rotate(line[:length-change], change)
	}
}
