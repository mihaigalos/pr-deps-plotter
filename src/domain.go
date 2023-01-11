package main

type PullRequest struct {
	Name  string
	Href  string
	State string
	Title string

	Dependencies []*PullRequest
}

