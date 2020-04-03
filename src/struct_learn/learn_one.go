package struct_learn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
	test   int
}

type Wheel struct {
	Circle
	Spokes int
}

func Struct_one() {
	var w = Wheel{
		Circle: Circle{Point{
			X: 1,
			Y: 2,
		}, 7, 1},
		Spokes: 0,
	}

	fmt.Printf("%#v\n", w)
	w.Y = 10

	fmt.Printf("%#v\n", w)
}

const IssuesURL = "https://api.github.com/search/issues"

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

	Body string // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err }
	resp.Body.Close()
	return &result, nil
}