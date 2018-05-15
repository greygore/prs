package main

import "testing"

func TestGetSearchResults(t *testing.T) {
	// For a basic sanity test, we use Github's "octocat", which seems to have some canned responses
	// A more robust test would use more accounts, and test the returned data more thoroughly
	// We might also choose to mock the response, rather than call the actual API
	users := []string{"octocat"}
	items := getSearchResults(users)
	if len(items) != 5 {
		t.Errorf("incorrect number of results: expected %d, got %d", 5, len(items))
	}
	if items[1].Title != "Add girlsnberry.md" {
		t.Errorf("incorrect title: expected '%s', got '%s'", "Add girlsnberry.md", items[1].Title)
	}
	if items[3].PullRequest.HTML != "https://github.com/violet-org/boysenberry-repo/pull/16" {
		t.Errorf("incorrect URL: expected '%s', got '%s'", "https://github.com/violet-org/boysenberry-repo/pull/16", items[3].PullRequest.HTML)
	}
}
