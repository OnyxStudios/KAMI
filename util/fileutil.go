package util

import (
	"fmt"
	"io/ioutil"
	"kami/resources"
)

func SReadAsset(path string) string {
	return string(ReadAsset(path))
}

func AssetExists(path string) bool {
	_, err := resources.AssetInfo(path)
	return err == nil
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

func CheckReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}



