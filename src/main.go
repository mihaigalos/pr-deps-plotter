package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	url := args[0]
	token := args[1]

	pr := Read(url, token)
	write(pr)
}
