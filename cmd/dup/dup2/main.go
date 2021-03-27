package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args) > 1 {
		for _, filename := range os.Args[1:] {
			f, err := os.Open(filepath.Clean(filename))
			if err != nil {
				log.Fatalln(err)
			}
			defer func() {
				err := f.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}()
			countLines(f, counts)
		}
	} else {
		countLines(os.Stdin, counts)
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Println(count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		log.Fatalln(err)
	}
}
