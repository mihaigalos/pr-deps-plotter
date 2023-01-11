package main

import (
	"fmt"
	"os"
	"text/template"
)

func write(prs []PullRequest){

	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("pr_open").Parse(pr_open)
	t.New("pr_merged").Parse(pr_merged)
	t.New("pr_closed").Parse(pr_closed)

	t.ExecuteTemplate(os.Stdout, "header", "")
	for _, pr := range prs {
		switch pr.State {
		case "open":
			t.ExecuteTemplate(os.Stdout, "pr_open", pr)
		case "merged":
			t.ExecuteTemplate(os.Stdout, "pr_merged", pr)
		case "closed":
			t.ExecuteTemplate(os.Stdout, "pr_closed", pr)
		}
	}
	for _, pr := range prs {
		if pr.Dependencies != nil {
			for _, dep := range pr.Dependencies {
				fmt.Println("    \"" + pr.Name + "\" -> \"" + dep.Name + "\"")
			}
		}
	}
	t.ExecuteTemplate(os.Stdout, "footer", "")
}

