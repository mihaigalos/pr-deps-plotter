package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestIntegrationWorks_whenTypical(t *testing.T) {
	expected := `digraph D {

    node [shape=plaintext]

    "wisespace-io/yubico-rs/26" [
    label=<
    <table border="0" cellborder="1" cellspacing="0" href="https://github.com/wisespace-io/yubico-rs/pull/26">
      <tr><td bgcolor="#4c278d"><font color="#ffffff">wisespace-io/yubico-rs/26</font></td></tr>
      <tr><td bgcolor="#ffffff"><font color="#000000">merged</font></td></tr>
      <tr><td bgcolor="#657a7f"><font color="#ffffff">Dockerize OTP validation functionality</font></td></tr>
    </table>>
    ];

    "wisespace-io/yubico-rs/25" [
    label=<
    <table border="0" cellborder="1" cellspacing="0" href="https://github.com/wisespace-io/yubico-rs/pull/25">
      <tr><td bgcolor="#4c278d"><font color="#ffffff">wisespace-io/yubico-rs/25</font></td></tr>
      <tr><td bgcolor="#ffffff"><font color="#000000">merged</font></td></tr>
      <tr><td bgcolor="#657a7f"><font color="#ffffff">Fix minimal example</font></td></tr>
    </table>>
    ];

    "wisespace-io/yubico-rs/26" -> "wisespace-io/yubico-rs/25"

}
`

	pr := Read("https://github.com/wisespace-io/yubico-rs/pull/26", "")

    r, w, _ := os.Pipe()
    os.Stdout = w
	Write(pr)
	w.Close()
	actual, _ := ioutil.ReadAll(r)

	expected_lines := strings.Split(expected, "\n")
	actual_lines := strings.Split(string(actual), "\n")

	if len(actual_lines) != len(expected_lines) {
		t.Errorf("len(actual) %d != len(expected) %d", len(actual_lines), len(expected_lines))
	}

	for i, actual_line := range actual_lines {
		if actual_line != expected_lines[i] {
			t.Errorf("No Match: %s != %s", actual_line, expected_lines[i])
		}
	}

}
