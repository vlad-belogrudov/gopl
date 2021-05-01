// Rotate string left or right a number of positions
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
	if shift < 0 {
		shift = length + shift
	}
	line = append(line, line[0:shift]...)[shift:]
	fmt.Println(string(line))
}
