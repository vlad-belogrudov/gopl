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

func search(terms []string) error {
	result, err := github.SearchIssues(terms)
	if err != nil {
		return err
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
			if time.Since(item.CreatedAt).Hours() > 24*30 {
				this_month = false
				color.Yellow("#### less than year ago ####")
			}
		}
		if this_year {
			if time.Since(item.CreatedAt).Hours() > 24*365 {
				this_year = false
				color.Red("#### more than year ago ####")
			}
		}
		fmt.Printf("#%-5d %s %20.20s %.55q\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}
	return nil
}

func getToken() (string, error) {
	const userPassVar = "GHUSERPASS"
	login, ok := os.LookupEnv(userPassVar)
	if !ok {
		return "", fmt.Errorf("%s is not set", userPassVar)
	}
	return login, nil
}

func create(repo, title, text string) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	issue, err := github.CreateIssue(token, repo, title, text)
	if err != nil {
		return err
	}
	fmt.Printf("Created issue %d at %s (state: %s)\n",
		issue.Number, issue.CreatedAt, issue.State)
	return nil
}

func help() {
	log.Fatalln("Usage: issues <create|update|close|search|show> [optinal terms]")
}

func main() {
	if len(os.Args) < 2 {
		help()
	}
	var err error
	switch os.Args[1] {
	case "search":
		err = search(os.Args[2:])
	case "create":
		if len(os.Args) != 5 {
			log.Fatalln("need additional arguments - full repo name, title, text")
		}
		err = create(os.Args[2], os.Args[3], os.Args[4])
	default:
		help()
	}
	if err != nil {
		log.Fatalln(err)
	}
}
