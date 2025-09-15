package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <sources> [flags]\n", os.Args[0])
		fmt.Printf("Type %s -h for help\n", os.Args[0])
		return
	}

	parseMode := 0 // 0 - sources file; 1 - output files name
	var sources []string
	var output string

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

	fmt.Println(sources)
	fmt.Println(output)
}