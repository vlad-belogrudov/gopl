package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type wstat struct {
	count int
	files map[string]struct{}
}

func main() {
	stats := make(map[string]*wstat)
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
			countLines(f, stats, filename)
		}
	} else {
		countLines(os.Stdin, stats, "stdin")
	}
	for line, st := range stats {
		if st.count > 1 {
			fmt.Printf("%dx %q found in:\n", st.count, line)
			for filename := range st.files {
				fmt.Println("\t", filename)
			}
		}
	}
}

func countLines(f *os.File, stats map[string]*wstat, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if stats[line] == nil {
			stats[line] = &wstat{
				files: make(map[string]struct{}),
				count: 0,
			}
		}
		stats[line].count++
		stats[line].files[filename] = struct{}{}
	}
	if err := input.Err(); err != nil {
		log.Fatalln(err)
	}
}
