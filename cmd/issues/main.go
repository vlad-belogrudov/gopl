// output table with results from Github search
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
		fmt.Printf("#%-5d %s %s %20.20s %.55q\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.State, item.User.Login, item.Title)
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

func show(repo string, number int) error {
	issue, err := github.GetIssue(repo, number)
	if err != nil {
		return err
	}
	fmt.Printf("Number:     %d\n", issue.Number)
	fmt.Printf("Title:      %s\n", issue.Title)
	fmt.Printf("Created at: %s\n", issue.CreatedAt)
	fmt.Printf("Created by: %s\n", issue.User.Login)
	fmt.Printf("State:      %s\n", issue.State)
	fmt.Printf("URL:        %s\n", issue.HTMLURL)
	fmt.Printf("Text:       %s\n", issue.Body)
	for _, comment := range issue.Comments {
		fmt.Println("### Comment:")
		fmt.Printf("\tCreated at: %s\n", comment.CreatedAt)
		fmt.Printf("\tCreated by: %s\n", comment.User.Login)
		fmt.Printf("\tText:       %s\n", comment.Body)
	}

	return nil
}

func close(repo string, number int) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	if err := github.CloseIssue(token, repo, number); err != nil {
		return err
	}
	fmt.Printf("Closed issue %d at %s\n", number, repo)
	return nil
}

func update(repo string, number int, title string) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	if err := github.UpdateIssue(token, repo, number, title); err != nil {
		return err
	}
	fmt.Printf("Updated issue %d at %s\n", number, repo)
	return nil
}

func comment(repo string, number int, text string) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	if err := github.CreateComment(token, repo, number, text); err != nil {
		return err
	}
	fmt.Printf("Commented issue %d at %s\n", number, repo)
	return nil
}

func help() {
	log.Fatalln("Usage: issues <create|update|comment|close|search|show> [optinal terms]")
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
	case "close":
		if len(os.Args) != 4 {
			log.Fatalln("need additional arguments - full repo name and issue number")
		}
		var number int
		number, err = strconv.Atoi(os.Args[3])
		if err != nil {
			break
		}
		err = close(os.Args[2], number)
	case "show":
		if len(os.Args) != 4 {
			log.Fatalln("need additional arguments - full repo name and issue number")
		}
		var number int
		number, err = strconv.Atoi(os.Args[3])
		if err != nil {
			break
		}
		err = show(os.Args[2], number)
	case "update":
		if len(os.Args) != 5 {
			log.Fatalln("need additional arguments - full repo name, issue number and new title")
		}
		var number int
		number, err = strconv.Atoi(os.Args[3])
		if err != nil {
			break
		}
		err = update(os.Args[2], number, os.Args[4])
	case "comment":
		if len(os.Args) != 5 {
			log.Fatalln("need additional arguments - full repo name, issue number and comment")
		}
		var number int
		number, err = strconv.Atoi(os.Args[3])
		if err != nil {
			break
		}
		err = comment(os.Args[2], number, os.Args[4])
	default:
		help()
	}
	if err != nil {
		log.Fatalln(err)
	}
}
