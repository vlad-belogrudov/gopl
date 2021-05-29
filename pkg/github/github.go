package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const searchIssuesUrl = "https://api.github.com/search/issues"
const createIssueUrl = "https://api.github.com/repos/%s/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type NewIssue struct {
	Title string `json:"title"`
	Text  string `json:"body"`
}

type NewIssueResult struct {
	Number    int       `json:"number"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues asks Github for terms
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	log.Println(q)
	resp, err := http.Get(searchIssuesUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed request: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateIssue creates new issue at Github repo
func CreateIssue(token, repo, title, text string) (*NewIssueResult, error) {
	issue := NewIssue{Title: title, Text: text}
	data, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf(createIssueUrl, repo), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed request: %s", resp.Status)
	}
	var result NewIssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
