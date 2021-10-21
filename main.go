package main

import (
	"encoding/json"
	"fmt"
	"log"

	kafka "github.com/brwillian/kafka-consumer-api/config"
	models "github.com/brwillian/kafka-consumer-api/models"
	services "github.com/brwillian/kafka-consumer-api/services"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	callServices()
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

		services.SaveDb(val)

	}
}
