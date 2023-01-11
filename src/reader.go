package main

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
func read() []PullRequest {
	pr_aim_182 := PullRequest{"aim/182", "https://github.com/mihaigalos/aim/pull/182", "closed", "Fix OpenSSL dep - revision 300.0.10+3.0.6 yanked", nil}
	pr_aim_189 := PullRequest{"aim/189", "https://github.com/mihaigalos/aim/pull/189", "merged", "Http serve: Propagate result", []*PullRequest{&pr_aim_182}}
	pr_aim_208 := PullRequest{"aim/208", "https://github.com/mihaigalos/aim/pull/208", "open", "Fix vulnerability CVE-2022-23639", []*PullRequest{&pr_aim_189}}

	prs := []PullRequest{pr_aim_208, pr_aim_189, pr_aim_182}

	return prs
}
