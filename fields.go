package imgstojson

import (
	"regexp"
	"strconv"
	"strings"
)

// ImgData is a struct of exif/iptc attributes we want from the exiftool
type ImgData struct {
	ID           string   `json:"ID,omitempty"`
	Caption      string   `json:"caption,omitempty"`
	Copyright    string   `json:"copyright,omitempty"`
	CreatedDate  string   `json:"createdDate,omitempty"`
	Description  string   `json:"description,omitempty"`
	Directory    string   `json:"directory,omitempty"`
	FNumber      string   `json:"fNum,omitempty"`
	Filename     string   `json:"filename,omitempty"`
	FocalLength  string   `json:"focalLength,omitempty"`
	ImageHeight  int64    `json:"imageHeight,omitempty"`
	ImageWidth   int64    `json:"imageWidth,omitempty"`
	ISO          int64    `json:"iso,omitempty"`
	Keywords     []string `json:"keywords,omitempty"`
	Lens         string   `json:"lens,omitempty"`
	ShutterSpeed string   `json:"shutterSpeed,omitempty"`
	Title        string   `json:"title,omitempty"`
}

var defaultFields = map[string]bool{
	"Caption-Abstract":            true,
	"Copyright":                   true,
	"Create Date":                 true,
	"Description":                 true,
	"Directory":                   true,
	"F Number":                    true,
	"File Name":                   true,
	"File Size":                   true,
	"Focal Length In 35mm Format": true,
	"ISO":           true,
	"Image Height":  true,
	"Image Number":  true,
	"Image Width":   true,
	"Keywords":      true,
	"Lens":          true,
	"Shutter Speed": true,
	"Subject":       true,
	"Title":         true,
}

// Transcribe takes an unformatted map of raw exifdata and returns the struct we actually want
func Transcribe(m map[string]string) ImgData {
	var re = regexp.MustCompile(`\D`)

	// about zero chance we'll two images created at the same time, use this as a unique ID
	id := re.ReplaceAllString(m["Create Date"], "")

	height, _ := strconv.ParseInt(m["Image Height"], 10, 64)
	width, _ := strconv.ParseInt(m["Image Width"], 10, 64)
	iso, _ := strconv.ParseInt(m["ISO"], 10, 64)

	return ImgData{
		id,
		m["Caption-Abstract"],
		m["Copyright"],
		m["Create Date"],
		m["Description"],
		m["Directory"],
		m["F Number"],
		m["File Name"],
		m["Focal Length In 35mm Format"],
		height,
		width,
		iso,
		strings.Split(m["Keywords"], ", "),
		m["Lens"],
		m["Shutter Speed"],
		m["Title"],
	}
}

// InDefaults returns a boolean if it appears in the default metadata map
func InDefaults(k string) bool {
	if defaultFields[k] == true {
		return true
	}

	return false
}
