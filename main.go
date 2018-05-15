package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

	items := getSearchResults(users)
	displayItems(items)
}

func endpoint(user string) string {
	q := url.Values{}
	q.Add("q", fmt.Sprintf("type:pr state:closed author:%s", user))

	u := url.URL{Scheme: "https", Host: "api.github.com", Path: "search/issues", RawQuery: q.Encode()}
	return u.String()
}

func getSearchResults(users []string) []searchItem {
	items := []searchItem{}

	for _, user := range users {
		resp, err := http.Get(endpoint(user))
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

	return items
}

func displayItems(items []searchItem) {
	var currUser string
	for _, item := range items {
		if item.User.Login != currUser {
			currUser = item.User.Login
			fmt.Println(item.User.Login)
		}
		fmt.Printf("\t%s\n\t\t%s\n", item.Title, item.PullRequest.HTML)
	}
}
