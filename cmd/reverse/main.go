package main

import (
	"fmt"
	"os"

	"github.com/vlad-belogrudov/gopl/pkg/reverse"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "need string to revert")
		os.Exit(1)
	}
	bytes := []byte(os.Args[1])
	if err := reverse.RevertUTF8Bytes(bytes); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
