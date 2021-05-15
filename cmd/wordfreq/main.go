// count word frequency
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type wordStat struct {
	word  string
	count int
}

func main() {
	wf := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wf[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
	}
	stats := make([]wordStat, 0, len(wf))
	for w, c := range wf {
		stats = append(stats, wordStat{word: w, count: c})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].count > stats[j].count
	})
	for _, s := range stats {
		fmt.Printf("%s\t%d\n", s.word, s.count)
	}
}
