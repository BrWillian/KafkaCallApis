package kafka

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	topic         = "FrotaBoiola"
	brokerAddress = "192.168.250.11:9092"
	groupId       = "Kafka-IA-3"
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
		"bootstrap.servers": brokerAddress,
		"group.id":          groupId,
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
