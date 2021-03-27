package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		log.Fatalln(err)
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Println(count, line)
		}
	}
}
