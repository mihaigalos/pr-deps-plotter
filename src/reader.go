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
    references := getReferences(base_pr_url, token)
	base_pr_name := splitBasePRName(base_pr_url)
    
	deps := []*PullRequest {}
    for _,ref := range references {
    	path := strings.Split(ref, "https://")[1]
		remainder := strings.Split(path, "/")
		name := strings.Join(remainder[1:], "/")
		fmt.Println("# read: ",name)
		state := getPRState(ref, "state", token)
		description := getPRInfo(ref, "title", token)
		dep := PullRequest{name, ref, state, description, nil}
    	deps = append(deps, &dep)
    }

	base_pr_state := getPRState(base_pr_url, "state", token)
	base_pr_description := getPRInfo(base_pr_url, "title", token)

	result := PullRequest {base_pr_name, base_pr_url, base_pr_state, base_pr_description, deps}

	return result
}

func getPRState(url string, field string, token string) string {
	fmt.Println("# getPRState: ",url)
	state := getPRInfo(url, "state", token)
	_ = getPRInfo(url, "merged", token)
	// is_merged, e := strconv.ParseBool(merged)
	if state == "closed"{
		state = "merged"
	}

	return state
}


func splitBasePRName(url string) string {
	url = prependApi(url)
	url = addReposToURLPath(url)
	url = addPullsToURLPath(url)
    split := strings.Split(url, "https://")[1]
	split = strings.Split(split, "/repos/")[1]
	split2 := strings.Split(split, "/pulls/")

	return split2[0] + "/" + split2[1]
}

func prependApi(url string) string {
	split := strings.Split(url, "://")
	schema := split[0]
	remainder := split[1]
	if !strings.HasPrefix(remainder, "api.") {
		url = schema + "://api." + remainder
	}

	return url
}

func addReposToURLPath(url string) string {
	split := strings.Split(url, "://")
	schema := split[0]
	remainder := split[1]

	split2 := strings.Split(remainder, "/")
	domain := split2[0]
	path := strings.Join(split2[1:],"/")
	if !strings.HasPrefix(path, "repos") {
		url = schema + "://" + domain + "/repos/" + path
	}

	return url
}

func addPullsToURLPath(url string) string {
	split := strings.Split(url, "/")
	path := split[0:len(split)-2]
	pr_number := split[len(split)-1]
	if !strings.HasSuffix(path[len(path)-1], "pulls") {
		url = strings.Join(path, "/") + "/pulls/" + pr_number
	}

	return url
}

func getPRInfo(url string, field string, token string) string {
	url = prependApi(url)
	url = addReposToURLPath(url)
	url = addPullsToURLPath(url)
	fmt.Println("# ++> ",url)

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

func getPRBody(url string, token string) []string {
	prInfo := getPRInfo(url, "body", token)
	return regexp.MustCompile("\r?\n").Split(prInfo, -1)
}

func getReferences(url string, token string) []string {
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
