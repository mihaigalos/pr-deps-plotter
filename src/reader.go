package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	req, _ := http.NewRequest("GET", "https://api.github.com/repos/scumjr/yubikeyedup/pulls/10", nil)
	// ...
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

	fmt.Println(objmap)
	

	fmt.Println("-------------------------")


    var prInfo string
    // Notice the dereferencing asterisk *
    err = json.Unmarshal([]byte(*objmap["body"]), &prInfo)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("%+v\n", prInfo)

	prs := []PullRequest{}

	return prs
}
