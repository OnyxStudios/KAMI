package util

import (
	"fmt"
	"io/ioutil"
	"kami/resources"
)

func SReadAsset(path string) string {
	return string(ReadAsset(path))
}

func ReadAsset(path string) []byte {
	data, err := resources.Asset(path)
	FCheckErr(err, fmt.Sprintf("asset file %v does not exist!", path))
	return data
}

func SReadFile(path string) string {
	return string(ReadFile(path))
}

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	FCheckErr(err, "unable to read file: %v")
	return data
}



