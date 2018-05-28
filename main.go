package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: make configurable
var defaultRootPath = "./images"

func walker(root string) []string {
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

func main() {
	imgData := GetMetadata(defaultRootPath)

	data, _ := json.Marshal(imgData)

	// TODO: make configuraable
	file, err := os.Create("out.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := string(data[:])
	fmt.Fprintf(file, s)
}
