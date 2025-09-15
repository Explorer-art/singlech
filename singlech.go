package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"path/filepath"
)

var sources []string
var output string
var implIncludes []string
var headerIncludes []string

func containsSlice(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

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

func ImplHandler(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(path + " not exists")
	}

	defer file.Close()

	var code string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplited := strings.Split(line, " ")

		if len(lineSplited) > 1 && lineSplited[0] == "#include" {
			if !containsSlice(implIncludes, line) {
				implIncludes = append(implIncludes, line)
			}
		} else if len(lineSplited) > 1 && lineSplited[0] == "#pragma" && lineSplited[1] == "once" {
			continue
		} else {
			code += line + "\n"
		}
	}

	return code
}

func HeaderHandler(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(path + " not exists")
	}

	defer file.Close()

	var code string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplited := strings.Split(line, " ")

		if len(lineSplited) > 1 && lineSplited[0] == "#include" {
			if !containsSlice(headerIncludes, line) {
				headerIncludes = append(headerIncludes, line)
			}
		} else if len(lineSplited) > 1 && lineSplited[0] == "#pragma" && lineSplited[1] == "once" {
			continue
		} else {
			code += line + "\n"
		}
	}

	return code
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

	// fmt.Println(sources)
	// fmt.Println(output)

	for _, file := range sources {
		fileExt := filepath.Ext(file)

		if fileExt == ".c" {
			fmt.Println(ImplHandler(file))
		} else if fileExt == ".h" {
			fmt.Println(HeaderHandler(file))
		}
	}

	// fmt.Println(includes)
}