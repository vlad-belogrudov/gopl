// Package reverse helps to revert things like strings or runes
package reverse

import (
	"fmt"
	"unicode/utf8"
)

func RevertString(line string) string {
	runes := []rune(line)
	RevertRunes(runes)
	return string(runes)
}

func RevertRunes(line []rune) {
	for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
		line[i], line[j] = line[j], line[i]
	}
}

func move(line []byte, num int) {
	if num == 0 {
		return
	} else if num < 0 {
		for i := 0; i < len(line)+num; i++ {
			line[i] = line[i-num]
		}
	} else {
		for i := len(line) - 1; i >= num; i-- {
			line[i] = line[i-num]
		}
	}
}

func RevertUTF8Bytes(line []byte) error {
	first, last := 0, len(line)
	for last-first > 1 {
		r1, s1 := utf8.DecodeRune(line[first:last])
		if r1 == utf8.RuneError && s1 == 1 {
			return fmt.Errorf("wrong start of utf8 %v", line[first:last])
		}
		r2, s2 := utf8.DecodeLastRune(line[first+s1 : last])
		if r2 == utf8.RuneError && s2 == 1 {
			return fmt.Errorf("wrong end of utf8 string %v", line[first:last])
		}
		if s2 == 0 {
			return nil
		}
		move(line[first:last], s2-s1)
		utf8.EncodeRune(line[last-s1:last], r1)
		utf8.EncodeRune(line[first:first+s2], r2)
		first += s2
		last -= s1
	}
	return nil
}
