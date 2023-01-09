// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
	"os"
	"text/template"
)

type PullRequestData struct {
	Name  string
	Href  string
	State string
	Title string
}

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

var footer = `
}
`

func main() {
	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("pr_open").Parse(pr_open)
	t.New("pr_merged").Parse(pr_merged)
	t.New("pr_closed").Parse(pr_closed)
	prs := []PullRequestData{
		{"aim/208", "https://github.com/mihaigalos/aim/pull/208", "open", "Fix vulnerability CVE-2022-23639"},
		{"aim/189", "https://github.com/mihaigalos/aim/pull/189", "merged", "Http serve: Propagate result"},
		{"aim/182", "https://github.com/mihaigalos/aim/pull/182", "closed", "Fix OpenSSL dep - revision 300.0.10+3.0.6 yanked"},
	}
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
	t.ExecuteTemplate(os.Stdout, "footer", "")
}
