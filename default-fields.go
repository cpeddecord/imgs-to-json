package main

var DefaultFields = [...]string{
	"Artist",
	"Camera Model Name",
	"Caption-Abstract",
	"Copyright",
	"Date Created",
	"Description",
	"Directory",
	"Exposure Time",
	"F Number",
	"File Name",
	"File Size",
	"Focal Length In 35mm Format",
	"ISO",
	"Image Description",
	"Image Height",
	"Image Number",
	"Image Size",
	"Image Width",
	"Keywords",
	"Label",
	"Lens",
	"Lens ID",
	"Lens Info",
	"Lens Model",
	"Object Name",
	"Original Document ID",
	"Shutter Speed",
	"Shutter Speed Value",
	"Subject",
	"Time Created",
	"Title",
	"X Resolution",
	"Y Resolution",
}

func InDefaults(k string) bool {
	for _, s := range DefaultFields {
		if s == k {
			return true
		}
	}

	return false
}
