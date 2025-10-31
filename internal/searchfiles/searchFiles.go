package searchfiles

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// SearchForFiles scans the given directory recursively
// and returns a list of file paths that match the provided
// extension and/or filename substring.
func SearchForFiles(dirPath string, fileExtension string, fileName string) ([]string, error) {
	var results []string
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		var matches bool = true
		if err != nil {
			return err
		}

		if d.IsDir() {
			matches = false
			return nil
		}

		if fileExtension != "" {
			matches = matches && filepath.Ext(path) == fileExtension
		}

		if fileName != "" {
			matches = matches && strings.Contains(d.Name(), fileName)
		}

		if matches {
			results = append(results, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return results, nil
	}
}
