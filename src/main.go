// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

func main() {
	prs := read("https://api.github.com/repos/scumjr/yubikeyedup/pulls/10")
	write(prs)
}
