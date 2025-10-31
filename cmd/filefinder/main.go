package main

import (
	"flag"
	"fmt"
	"go-file-finder/internal/searchfiles"
	"os"
)

func main() {
	var dirPath = flag.String("path", "", "Path to the target directory (required). Example: -path C:\\Games")
	var fileExtension = flag.String("ext", "", "Filter by file extension (optional). Example: -ext .exe")
	var fileName = flag.String("name", "", "Filter by filename substring (optional). Example: -name fullHD")
	flag.Parse()

	if *dirPath == "" {
		fmt.Fprintln(os.Stderr, "-path flag is required. Usage: go run main.go -path <file.txt>")
		os.Exit(1)
	}

	results, err := searchfiles.SearchForFiles(*dirPath, *fileExtension, *fileName)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Search failed: %v", err)
		os.Exit(1)
	}

	if len(results) == 0 {
		fmt.Println("No matching files found.")
		return
	}

	for _, file := range results {
		fmt.Println("Found matching file:", file)
	}
}
