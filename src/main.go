package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	fmt.Println(args)
	url := args[0]
	token := args[1]

	pr := Read(url, token)
	Write(pr)
}
