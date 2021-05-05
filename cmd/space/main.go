package main

import (
	"fmt"
	"os"

	"github.com/vlad-belogrudov/gopl/pkg/space"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, `need word to brush, e.g "hello   bye end"`)
		os.Exit(1)
	}
	fmt.Println(string(space.Brush([]byte(os.Args[1]))))
}
