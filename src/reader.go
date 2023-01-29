package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"strconv"
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

func read(base_pr_url string, token string) PullRequest {
    references := getPRReferences(base_pr_url, token)
    
	deps := []*PullRequest {}
    for _,ref := range references {
		dep := read(ref, token)
		deps = append(deps, &dep)
    }

	base_pr_name := splitBasePRName(base_pr_url)
	base_pr_state := getPRState(base_pr_url, "state", token)
	base_pr_description := getPRInfo(base_pr_url, "title", token)

	result := PullRequest {base_pr_name, base_pr_url, base_pr_state, base_pr_description, deps}

	return result
}

func getPRState(url string, field string, token string) string {
	state := getPRInfo(url, "state", token)
	merged := getPRInfo(url, "merged", token)
	is_merged, _ := strconv.ParseBool(merged)

	if state == "closed" && is_merged {
		state = "merged"
	}

	return state
}

func getPRInfo(url string, field string, token string) string {
	url = prependApi(url)
	url = addReposToURLPath(url)
	url = addPullsToURLPath(url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Authorization", token)
	resp, _ := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal(body, &objmap); err != nil {
		log.Fatal(err)
	}

	return unmarshal(objmap, field);
}

func getPRBody(url string, token string) []string {
	prInfo := getPRInfo(url, "body", token)
	return regexp.MustCompile("\r?\n").Split(prInfo, -1)
}

func getPRReferences(url string, token string) []string {
	prInfo := getPRBody(url, token)
	result := []string{}
	for _, line := range prInfo {
		if strings.HasSuffix(line, ".") {
			line = line[:len(line)-1]
		}
		r := regexp.MustCompile("^[dD]epends-[oO]n: (.*)+")
		res := r.FindStringSubmatch(line)
		if len(res) > 0 {
			result = append(result, res[1])
		}
	}
	return result
}

func unmarshal(objmap map[string]*json.RawMessage, field string) string {
	var prInfoString string
	var prInfoBool bool

	err := json.Unmarshal([]byte(*objmap[field]), &prInfoString)

	if err != nil {
		err = json.Unmarshal([]byte(*objmap[field]), &prInfoBool)
		if err == nil {
			return strconv.FormatBool(prInfoBool);
		} else {
			log.Fatal(err)
		}

	}
	return prInfoString;
}

