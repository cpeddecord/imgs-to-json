package main

import (
	"os"
	"path/filepath"
	"strings"
)

func Walker(root string) []string {
	var paths []string

	walk := func(path string, info os.FileInfo, err error) error {
		// TODO make configurable
		// TODO have default file types, jpg, jpeg, gif, etc.
		if strings.Contains(path, ".jpg") {
			paths = append(paths, path)
		}
		return nil
	}

	err := filepath.Walk(root, walk)
	if err != nil {
		panic(err)
	}

	return paths
}
