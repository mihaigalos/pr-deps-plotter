package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestReadWorks_whenTypical(t *testing.T) {
	expected := "https://github.com/scumjr/yubikeyedup/pull/9"
	actual := ""
	prInfo := GetPRInfo("https://api.github.com/repos/scumjr/yubikeyedup/pulls/10")

	for _,line := range prInfo {
		if strings.HasSuffix(line, ".") {
			line = line[:len(line)-1]
		}
		r := regexp.MustCompile("^[dD]epends-[oO]n: (.*)+")
		res := r.FindStringSubmatch(line)
		if len(res) > 0 {
			actual = res[1]
		}
	}

	if actual != expected {
		t.Errorf("No Match: %s != %s", actual, expected)
	}
}
