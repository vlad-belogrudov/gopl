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
	length := len(line)
	shift %= length
	next := 0
	temp := line[next]
	for range line {
		next = rotated(next, shift, length)
		fmt.Printf("next %d stores %c, line %s\n", next, temp, string(line))
		temp, line[next] = line[next], temp
	}
	fmt.Println(string(line))
}

func rotated(pos, shift, length int) int {
	pos -= shift
	if pos < 0 {
		pos = length + pos
	}
	return pos % length
}
