package main

import (
    "os"
)

func main() {
    args := os.Args[1:]
	pr := args[0]
	token := args[1]

	prs := read(pr, token)
	write(prs)
}
