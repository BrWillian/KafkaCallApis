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
	capacete_api      = "http://127.0.0.1:5002/api/1.0/capacete"
	classificador_api = "http://192.168.250.110:5003/api/classificador/veicular"
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
func ConsumeCapaceteApi(image string) []byte {
	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", capacete_api, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err.Error())
	}
	data, _ := ioutil.ReadAll(response.Body)

	return data
}

func ConsumeClassificadorApi(image string) []byte {
	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", classificador_api, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err.Error())
	}
	data, _ := ioutil.ReadAll(response.Body)

	return data
}
