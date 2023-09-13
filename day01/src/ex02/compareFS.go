package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getSet(filename string) map[string]bool {
	file := openFile(filename)
	set := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		set[text] = true
	}
	return set
}

func openFile(filename string) *os.File {
	if !strings.HasSuffix(filename, ".txt") {
		fmt.Fprintf(os.Stderr, "Error: wrong file extension: \"%s\":\n", filename)
		os.Exit(1)
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return file
}

func compareFiles(oldFile string, newFile string) {
	oldSet := getSet(oldFile)
	newSet := getSet(newFile)

	for item := range oldSet {
		if _, ok := newSet[item]; !ok {
			fmt.Printf("REMOVED %s\n", item)
		}
	}
	for item := range newSet {
		if _, ok := oldSet[item]; !ok {
			fmt.Printf("ADDED %s\n", item)
		}
	}
}

func main() {
	oldFile := flag.String("old", "", "use old file")
	newFile := flag.String("new", "", "use new file")
	flag.Parse()

	if *oldFile != "" && *newFile != "" {
		compareFiles(*oldFile, *newFile)
	} else {
		fmt.Println("Use '--old' and '--new' flags to pass arguments")
	}
}