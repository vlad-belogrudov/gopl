// Package xkcd helps to find comics by keywords
package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	stemming "github.com/kljensen/snowball/english"
)

const XKCDURL = "https://xkcd.com/%d/info.0.json"

type database struct {
	Documents []string
	Index     map[string][]int
}

var db database

type Entry struct {
	number int
	Title  string `json:"title"`
	Text   string `json:"transcript"`
	URL    string `json:"img"`
}

type EntryStorage struct {
	entries []*Entry
	locker  sync.Mutex
}

func (s *EntryStorage) addEntry(e *Entry) {
	s.locker.Lock()
	total := len(s.entries)
	if total > 0 && total%100 == 0 {
		fmt.Printf("Got %d entries\n", len(s.entries))
	}
	defer s.locker.Unlock()
	s.entries = append(s.entries, e)
}

func download(wg *sync.WaitGroup, lastDigit int, store *EntryStorage) {
	defer wg.Done()
	fmt.Printf("downloading digit %d\n", lastDigit)
	for next := lastDigit; true; next += 10 {
		var entry Entry
		entry.number = next
		if next == 0 || next == 404 {
			store.addEntry(&entry)
			continue
		}
		resp, err := http.Get(fmt.Sprintf(XKCDURL, next))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get URL: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Could not get %d, finishing\n", next)
			return
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Bad status: %s\n", resp.Status)
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read response: %v\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(data, &entry); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse json: %v\n", err)
			os.Exit(1)
		}
		store.addEntry(&entry)
	}
}

var dropWords = map[string]bool{
	"a": true, "and": true, "be": true, "from": true, "have": true, "i": true,
	"in": true, "of": true, "that": true, "the": true, "to": true, "was": true,
	"were": true,
}

func cleanWords(line string) []string {
	// remove special chars
	line = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return ' '
	}, line)
	// remove frequent words and variations
	wordSet := make(map[string]bool)
	for _, w := range strings.Fields(line) {
		if !dropWords[w] {
			// normalize
			w = stemming.Stem(w, false)
			wordSet[strings.ToLower(w)] = true
		}
	}
	var words []string
	for w := range wordSet {
		words = append(words, w)
	}
	return words
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get home directory: %v\n", err)
		os.Exit(1)
	}
	datapath := path.Join(home, "xkcd.db")
	if _, err := os.Stat(datapath); os.IsNotExist(err) {
		fmt.Println("Downloading database...")
		startTime := time.Now()
		var wg sync.WaitGroup
		var store EntryStorage
		for digit := 0; digit < 10; digit++ {
			wg.Add(1)
			go download(&wg, digit, &store)
		}
		wg.Wait()
		duration := time.Since(startTime)
		fmt.Printf("Finished download in %d sec\n", duration.Milliseconds()/1000)

		sort.Slice(store.entries, func(i, j int) bool {
			return store.entries[i].number < store.entries[j].number
		})
		db.Index = make(map[string][]int)
		for _, e := range store.entries {
			db.Documents = append(db.Documents, e.URL)
			words := cleanWords(e.Text + e.Title)
			for _, w := range words {
				db.Index[w] = append(db.Index[w], e.number)
			}
		}

		data, err := json.Marshal(db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot encode database: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Write database: %d entries\n", len(db.Documents))
		if err := os.WriteFile(datapath, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot write to database: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Reading database from disk: %s\n", datapath)
		data, err := ioutil.ReadFile(datapath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read db: %v\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(data, &db); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot decode db: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Read %d entries\n", len(db.Documents))
	}
}

func Search(tokens []string) []string {
	if len(tokens) == 0 {
		return []string{}
	}
	matcher := make(map[int]int) // see how many tokens each document has
	for _, token := range tokens {
		token = stemming.Stem(token, false)
		for _, index := range db.Index[token] {
			matcher[index]++
		}
	}
	var results []string
	for index, found := range matcher {
		// if document has all tokens - take it
		if found == len(tokens) {
			results = append(results, db.Documents[index])
		}
	}
	return results
}
