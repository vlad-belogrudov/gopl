package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		for _, line := range strings.Split(string(content), "\n") {
			counts[line]++
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Println(count, line)
		}
	}
}
