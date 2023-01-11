// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t, _ := template.New("header").Parse(header)
	t.New("footer").Parse(footer)
	t.New("pr_open").Parse(pr_open)
	t.New("pr_merged").Parse(pr_merged)
	t.New("pr_closed").Parse(pr_closed)
	pr_aim_182 := PullRequest{"aim/182", "https://github.com/mihaigalos/aim/pull/182", "closed", "Fix OpenSSL dep - revision 300.0.10+3.0.6 yanked", nil}
	pr_aim_189 := PullRequest{"aim/189", "https://github.com/mihaigalos/aim/pull/189", "merged", "Http serve: Propagate result", []*PullRequest{&pr_aim_182}}
	pr_aim_208 := PullRequest{"aim/208", "https://github.com/mihaigalos/aim/pull/208", "open", "Fix vulnerability CVE-2022-23639", []*PullRequest{&pr_aim_189}}

	prs := []PullRequest{pr_aim_208, pr_aim_189, pr_aim_182}
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
