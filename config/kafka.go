package config

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	topic   = "flautim_canonico"
	groupId = "Jacu-Estalo-Ia"
	offset  = "latest"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}
func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          groupId,
		"auto.offset.reset": offset,
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Erro ao consumir a mensagem: " + err.Error())
	}
	topics := []string{topic}
	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer foi iniciado!")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
