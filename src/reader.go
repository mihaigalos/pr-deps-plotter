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
