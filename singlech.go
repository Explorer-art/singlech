package main

import (
	"fmt"
	"os"
)

var sources []string
var output string

func parseArgs() {
	parseMode := 0 // 0 - sources file; 1 - output files name

	for _, arg := range os.Args[1:] {
		if arg == "-o" {
			parseMode = 1
			continue
		}

		if parseMode == 0 {
			sources = append(sources, arg)
		} else if parseMode == 1 {
			output = arg
			parseMode = 0
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <sources> [flags]\n", os.Args[0])
		fmt.Printf("Type %s -h for help\n", os.Args[0])
		return
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("Usage: %s <sources> [flags]\n\n", os.Args[0])
		fmt.Printf("Flags:\n-o --output <name>: output files name\n-h --help: help for flags\n")
		return
	}

	parseArgs()

	fmt.Println(sources)
	fmt.Println(output)
}