// Find comics by keywords
package main

import (
	"fmt"
	"os"

	"github.com/vlad-belogrudov/gopl/pkg/xkcd"
)

func main() {
	entries := xkcd.Search(os.Args[1:])
	for entry := range entries {
		fmt.Println(entry)
	}
}
