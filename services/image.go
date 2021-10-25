package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func ReadImage(imgPath string) string {
	bytes, err := ioutil.ReadFile(imgPath)

	if err != nil {
		fmt.Println(err.Error())
	}

	return toBase64(bytes)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
