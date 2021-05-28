// output table with results from Github search
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/fatih/color"
	"github.com/vlad-belogrudov/gopl/pkg/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d titles:\n", result.TotalCount)
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})
	this_month := true
	this_year := true
	color.Green("#### less than month ago ####")
	for _, item := range result.Items {
		if this_month {
			if time.Now().Sub(item.CreatedAt).Hours() > 24*30 {
				this_month = false
				color.Yellow("#### less than year ago ####")
			}
		}
		if this_year {
			if time.Now().Sub(item.CreatedAt).Hours() > 24*365 {
				this_year = false
				color.Red("#### more than year ago ####")
			}
		}
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}
}
