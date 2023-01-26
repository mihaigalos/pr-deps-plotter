package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)
//curl -s \
//  -H "Accept: application/vnd.github+json" \
//  -H "X-GitHub-Api-Version: 2022-11-28" \
//  https://api.github.com/repos/scumjr/yubikeyedup/pulls/10 | jq .body | sed -e "s|\\\r\\\n|\n|g"

var client = &http.Client{Timeout: 10 * time.Second}
type PrInfo struct {
    Body string
}

func read() []PullRequest {
	//pr_aim_182 := PullRequest{"aim/182", "https://github.com/mihaigalos/aim/pull/182", "closed", "Fix OpenSSL dep - revision 300.0.10+3.0.6 yanked", nil}
	//pr_aim_189 := PullRequest{"aim/189", "https://github.com/mihaigalos/aim/pull/189", "merged", "Http serve: Propagate result", []*PullRequest{&pr_aim_182}}
	//pr_aim_208 := PullRequest{"aim/208", "https://github.com/mihaigalos/aim/pull/208", "open", "Fix vulnerability CVE-2022-23639", []*PullRequest{&pr_aim_189}}

	//prs := []PullRequest{pr_aim_208, pr_aim_189, pr_aim_182}


	prs := []PullRequest{}

	return prs
}

func prependApi(url string) string {
	split := strings.Split(url, "://")
	schema := split[0]
	remainder := split[1]
	if !strings.HasPrefix(remainder, "api.") {
		url = schema+"://api." + remainder
    }

	return url
}

func GetPRInfo(url string) []string {
	url = prependApi(url)

	req, _ := http.NewRequest("GET", url , nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	resp, _ := client.Do(req)
	
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        panic(err.Error())
    }

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal(body, &objmap); err != nil {
		log.Fatal(err)
	}

    var prInfo string
    err = json.Unmarshal([]byte(*objmap["body"]), &prInfo)
    if err != nil {
        log.Fatal(err)
    }

	return regexp.MustCompile("\r?\n").Split(prInfo, -1)
}
