package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	kafka "github.com/brwillian/kafka-consumer-api/config"
	models "github.com/brwillian/kafka-consumer-api/models"
	routers "github.com/brwillian/kafka-consumer-api/routers"
	services "github.com/brwillian/kafka-consumer-api/services"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

func main() {
	go func() {
		callServices()
	}()
	handleRequests()
}
func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/version", routers.GetVersion)
	myRouter.HandleFunc("/api/health", routers.HealthCheck)
	myRouter.HandleFunc("/api/ready", routers.Ready)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func callServices() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	var kmsg models.KafkaMessage

	for msg := range msgChan {
		err := json.Unmarshal(msg.Value, &kmsg)
		if err != nil {
			fmt.Println(err.Error())
		}

		log.Println(string(msg.Value))
		val := services.GetResult(kmsg)

		go services.SaveDb(val)

	}
}
