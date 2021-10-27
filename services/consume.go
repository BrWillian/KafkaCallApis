package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ConsumeOcrApi(image string) []byte {

	jsonData := map[string]string{"image": image}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", os.Getenv("APIOCR_URL"), bytes.NewBuffer(jsonValue))

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

	request, _ := http.NewRequest("POST", os.Getenv("APICAPACETE_URL"), bytes.NewBuffer(jsonValue))

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

	request, _ := http.NewRequest("POST", os.Getenv("APICLASSIFICADOR_URL"), bytes.NewBuffer(jsonValue))

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
