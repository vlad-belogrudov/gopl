package main

import (
	"strings"
	"testing"
)

var input = []string{"hello", "bye", "asdf", "1234567890", "-1", "-2"}

func echo1(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	return strings.Join(args, " ")
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = echo1(input)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = echo2(input)
	}
}
