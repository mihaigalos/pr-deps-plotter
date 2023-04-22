package main

import (
	"bytes"
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

func writePR(pr PullRequest, t *template.Template) string {
	buf := &bytes.Buffer{}
	switch pr.State {
	case "open":
		t.ExecuteTemplate(buf, "pr_open", pr)
	case "merged":
		t.ExecuteTemplate(buf, "pr_merged", pr)
	case "closed":
		t.ExecuteTemplate(buf, "pr_closed", pr)
	}
	result := buf.String()
	if pr.Dependencies != nil {
		for _, dep := range pr.Dependencies {
			result += writePR(*dep, t)
			result += "    \"" + pr.Name + "\" -> \"" + dep.Name + "\""
		}
	}
	result += "\n"
	return result
}

func buildTemplate() *template.Template {
	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("pr_open").Parse(pr_open)
	t.New("pr_merged").Parse(pr_merged)
	t.New("pr_closed").Parse(pr_closed)
	return t
}

func Write(pr PullRequest) string {
	t := buildTemplate()
	buf := &bytes.Buffer{}
	t.ExecuteTemplate(buf, "header", "")
	result := buf.String()
	result += writePR(pr, t)

	buf = &bytes.Buffer{}
	t.ExecuteTemplate(buf, "footer", "")
	result += buf.String()
	return result
}
