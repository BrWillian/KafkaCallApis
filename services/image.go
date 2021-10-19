package service

import (
	"encoding/base64"
	"io/ioutil"
	"log"
)

func ReadImage(imgPath string) string {
	bytes, err := ioutil.ReadFile(imgPath)

	if err != nil {
		log.Fatal(err)
	}

	return toBase64(bytes)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
