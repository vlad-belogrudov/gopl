// Package space helps remove unnecessary spaces in byte slice
package space

import "unicode"

func Brush(input []byte) []byte {
	i := 0
	seenSpace := false
	for _, b := range input {
		if unicode.IsSpace(rune(b)) {
			if seenSpace {
				continue
			}
			b = ' '
			seenSpace = true
		} else {
			seenSpace = false
		}
		input[i] = b
		i++
	}
	return input[:i]
}
