package main

import (
	"strings"
)

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

func splitBasePRName(url string) string {
	url = prependApi(url)
	url = addReposToURLPath(url)
	url = addPullsToURLPath(url)
    split := strings.Split(url, "https://")[1]
	split = strings.Split(split, "/repos/")[1]
	split2 := strings.Split(split, "/pulls/")

	return split2[0] + "/" + split2[1]
}

