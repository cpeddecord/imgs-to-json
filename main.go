package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO: make configurable
var ROOT = "./images"

func main() {
	paths := Walker(ROOT)
	var imgData []map[string]string

	for _, p := range paths {
		imgData = append(imgData, Exif(p))
	}

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
