package main

import (
	"fmt"
	"os"
	"strings"
	"terminal_ai/cli"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(1)
	}

	if args[0] == "configure" {
		error := cli.Configure()
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
		return
	}

	prompt := strings.Join(args, " ")
	response, error := cli.Ask(prompt)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	fmt.Println(response)
}
