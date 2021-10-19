package main

import (
	"encoding/json"
	"fmt"

	kafka "github.com/brwillian/kafka-consumer-api/config"
	models "github.com/brwillian/kafka-consumer-api/models"
	services "github.com/brwillian/kafka-consumer-api/services"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaMessage struct {
	caminhoImagem string
}

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

		result := services.ConsumeOcrApi(services.ReadImage(kmsg.CaminhoImagem))

		fmt.Println(string(result))

	}
}
