// Separate each 3 digits
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "please supply a number")
		os.Exit(1)
	}
	number := os.Args[1]
	if _, err := strconv.Atoi(number); err != nil {
		fmt.Fprintf(os.Stderr, "not an integer: %v\n", err)
		os.Exit(1)
	}
	var buffer bytes.Buffer
	shift := len(number) % 3
	start := 0
	if number[0] == '-' || number[0] == '+' {
		start = 1
	}
	for i := 0; i < len(number); i++ {
		if i > start && i%3 == shift {
			buffer.WriteByte('\'')
		}
		buffer.WriteByte(number[i])
	}
	fmt.Println(buffer.String())
}
