package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var dirPath = flag.String("path", "", "Path to the target directory (required). Example: -path C:\\Games")
	var fileExtension = flag.String("ext", "", "Filter by file extension (optional). Example: -ext .exe")
	var fileName = flag.String("name", "", "Filter by filename substring (optional). Example: -name fullHD")
	flag.Parse()

	filepath.WalkDir(*dirPath, func(path string, d fs.DirEntry, err error) error {
		var matches bool = true
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if d.IsDir() {
			matches = false
			return nil
		}

		if *fileExtension != "" {
			matches = matches && filepath.Ext(path) == *fileExtension
		}

		if *fileName != "" {
			matches = matches && strings.Contains(d.Name(), *fileName)
		}

		if matches {
			fmt.Println("Found matching file:", d.Name())
		}

		return nil
	})
}
