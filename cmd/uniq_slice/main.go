// Uniq removes adjacent duplicates from []string
package main

import (
	"fmt"
	"os"

	"github.com/vlad-belogrudov/gopl/pkg/uniq"
)

func main() {
	fmt.Println(uniq.Uniq(os.Args[1:]))
}
