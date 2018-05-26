package main

import (
	"os/exec"
	"strings"
)

var exiftool = "exiftool"

func Exif(p string) map[string]string {
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

		if len(splits) == 1 {
			continue
		}

		k := strings.Trim(splits[0], " ")
		if defaultKey := InDefaults(k); defaultKey == false {
			continue
		}

		exif[k] = strings.Trim(splits[1], " ")
	}

	return exif
}
