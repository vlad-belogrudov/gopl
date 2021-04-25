// Separate each 3 digits, e.g. go run main.go 123456.1234567 -> 123'456.123'456'7
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "please supply a number")
		os.Exit(1)
	}
	number := os.Args[1]
	if _, err := strconv.ParseFloat(number, 64); err != nil {
		fmt.Fprintf(os.Stderr, "not a number: %v\n", err)
		os.Exit(1)
	}
	var buffer bytes.Buffer
	point := strings.IndexByte(number, '.')
	if -1 == point {
		point = len(number)
	}
	shift := point % 3
	start := 0
	if number[0] == '-' || number[0] == '+' {
		start = 1
	}
	for i := 0; i < len(number); i++ {
		if i == point {
			shift = (shift + 1) % 3
			start = i + 1
		} else if i > start && i%3 == shift {
			buffer.WriteByte('\'')
		}
		buffer.WriteByte(number[i])
	}
	fmt.Println(buffer.String())
}
