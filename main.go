package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO: make configurable
var ROOT = "./images"

// TODO: https://blog.golang.org/pipelines
func main() {
	paths := Walker(ROOT)
	var imgData []map[string]string

	c := make(chan map[string]string)

	for _, p := range paths {
		go Exif(p, c)
	}

	for range paths {
		imgData = append(imgData, <-c)
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
