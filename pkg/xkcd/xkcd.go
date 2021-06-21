// Package xkcd helps to find comics by keywords
package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const XKCDURL = "https://xkcd.com/%d/info.0.json"

type database struct {
	Documents []string
	Index     map[string][]int
}

var db database

type Entry struct {
	Title string `json:"title"`
	Text  string `json:"transcript"`
	URL   string `json:"img"`
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get home directory: %v\n", err)
		os.Exit(1)
	}
	datapath := path.Join(home, "xkcd.db")
	if _, err := os.Stat(datapath); os.IsNotExist(err) {
		db.Index = make(map[string][]int)
		db.Documents = []string{""}
		for next := 1; true; next++ {
			resp, err := http.Get(fmt.Sprintf(XKCDURL, next))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not get URL: %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusNotFound {
				break
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
			var entry Entry
			if err := json.Unmarshal(data, &entry); err != nil {
				fmt.Fprintf(os.Stderr, "Cannot parse json: %v\n", err)
				os.Exit(1)
			}
			db.Documents = append(db.Documents, entry.URL)
		}
		data, err := json.Marshal(db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot encode database: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Write database: %d entries", len(db.Documents)-1)
		if err := os.WriteFile(datapath, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot write to database: %v\n", err)
			os.Exit(1)
		}

	} else {
		data, err := ioutil.ReadFile(datapath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read db: %v\n", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(data, &db); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot decode db: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Read database: %d entries\n", len(db.Documents)-1)
	}
}

func Search(tokens []string) []string {

	return []string{}
}
