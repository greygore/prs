package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type searchItem struct {
	Title string `json:"title"`
	User  struct {
		Login string `json:"login"`
	}
	PullRequest struct {
		HTML string `json:"html_url"`
	} `json:"pull_request"`
}

type searchResults struct {
	Incomplete bool         `json:"incomplete_results"`
	Items      []searchItem `json:"items"`
	Total      int          `json:"total_count"`
}

func main() {
	users := os.Args[1:]
	items := []searchItem{}

	for _, user := range users {
		resp, err := http.Get(fmt.Sprintf("https://api.github.com/search/issues?q=type:pr%%20state:closed%%20author:%s", user))
		if err != nil {
			log.Printf("Unable to GET pull requests for %s", user)
		}
		defer resp.Body.Close()

		var results searchResults
		if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
			log.Printf("Unable to decode pull requests for %s", user)
		}
		items = append(items, results.Items...)
	}

	var currUser string
	for _, item := range items {
		if item.User.Login != currUser {
			currUser = item.User.Login
			fmt.Printf("%s:\n", item.User.Login)
		}
		fmt.Printf("\t%s\n\t\t%s\n", item.Title, item.PullRequest.HTML)
	}
}
