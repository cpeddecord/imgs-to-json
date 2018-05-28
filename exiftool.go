package main

import (
	"os/exec"
	"strings"
)

var exiftool = "exiftool"

// Exif runs the exiftool script and harvests wanted metadata from stdout
func Exif(p string, c chan map[string]string) {
	out, _ := exec.Command(exiftool, p).Output()
	str := string(out[:])
	arr := strings.Split(str, "\n")

	exif := make(map[string]string)

	for _, s := range arr {
		// EOF line break
		if s == "" {
			continue
		}

		splits := strings.Split(s, " : ")

		// some fields won't split correctly and we don't need those anyhow
		if len(splits) == 1 {
			continue
		}

		k := strings.Trim(splits[0], " ")
		if defaultKey := InDefaults(k); defaultKey == false {
			continue
		}

		exif[k] = strings.Trim(splits[1], " ")
	}

	c <- exif
}

// GetMetadata returns a slice of structs for wanted metadata given a root directory to walk from
// TODO: https://blog.golang.org/pipelines
func GetMetadata(root string) []ImgData {
	paths := walker(root)
	var imgDatas []ImgData

	c := make(chan map[string]string)

	for _, p := range paths {
		go Exif(p, c)
	}

	for range paths {
		imgStruct := Transcribe(<-c)
		imgDatas = append(imgDatas, imgStruct)
	}

	return imgDatas
}
