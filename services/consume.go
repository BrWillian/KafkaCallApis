package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ocr_api           = "http://192.168.250.110:5001/api/ocr"
	capacete_api      = "http://127.0.0.1:5003/api/1.0/capacete"
	classificador_api = "http://192.168.250.110:5002/api/classificador/veicular"
)

func ConsumeOcrApi(image string) []byte {

	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", ocr_api, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		return data
	}
	return nil

}
func ConsumeCapaceteApi(image string) []byte {
	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", capacete_api, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		return data
	}
	return nil
}

func ConsumeClassificadorApi(image string) []byte {
	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", classificador_api, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		return data
	}
	return nil
}
