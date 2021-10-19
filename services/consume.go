package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ocr_api           = "http://192.168.250.110:5001/api/ocr"
	capacete_api      = ""
	classificador_api = ""
)

func ConsumeOcrApi(image string) []byte {

	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", ocr_api, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err.Error())
	}
	data, _ := ioutil.ReadAll(response.Body)

	return data
}
