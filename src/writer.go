package main

import (
	"fmt"
	"os"
	"text/template"
)

var header = `digraph D {

    node [shape=plaintext]
`

var pr_open = `
    "{{ .Name }}" [
    label=<
    <table border="0" cellborder="1" cellspacing="0" href="{{ .Href }}">
      <tr><td bgcolor="#008000"><font color="#ffffff">{{ .Name }}</font></td></tr>
      <tr><td bgcolor="#ffffff"><font color="#000000">{{ .State }}</font></td></tr>
      <tr><td bgcolor="#657a7f"><font color="#ffffff">{{ .Title }}</font></td></tr>
    </table>>
    ];
`
var pr_merged = `
    "{{ .Name }}" [
    label=<
    <table border="0" cellborder="1" cellspacing="0" href="{{ .Href }}">
      <tr><td bgcolor="#4c278d"><font color="#ffffff">{{ .Name }}</font></td></tr>
      <tr><td bgcolor="#ffffff"><font color="#000000">{{ .State }}</font></td></tr>
      <tr><td bgcolor="#657a7f"><font color="#ffffff">{{ .Title }}</font></td></tr>
    </table>>
    ];
`
var pr_closed = `
    "{{ .Name }}" [
    label=<
    <table border="0" cellborder="1" cellspacing="0" href="{{ .Href }}">
      <tr><td bgcolor="#a02f2b"><font color="#ffffff">{{ .Name }}</font></td></tr>
      <tr><td bgcolor="#ffffff"><font color="#000000">{{ .State }}</font></td></tr>
      <tr><td bgcolor="#657a7f"><font color="#ffffff">{{ .Title }}</font></td></tr>
    </table>>
    ];
`

var footer = `}
`

func writePR(pr PullRequest, t *template.Template) {
	switch pr.State {
	case "open":
		t.ExecuteTemplate(os.Stdout, "pr_open", pr)
	case "merged":
		t.ExecuteTemplate(os.Stdout, "pr_merged", pr)
	case "closed":
		t.ExecuteTemplate(os.Stdout, "pr_closed", pr)
	}
	if pr.Dependencies != nil {
		for _, dep := range pr.Dependencies {
			writePR(*dep, t)
			fmt.Println("    \"" + pr.Name + "\" -> \"" + dep.Name + "\"")
		}
	}
	fmt.Println("")
}

func write(pr PullRequest) {
	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("pr_open").Parse(pr_open)
	t.New("pr_merged").Parse(pr_merged)
	t.New("pr_closed").Parse(pr_closed)

	t.ExecuteTemplate(os.Stdout, "header", "")
	writePR(pr, t)
	t.ExecuteTemplate(os.Stdout, "footer", "")
}
