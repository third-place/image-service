package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Producer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}

func CreateProducer() Producer {
	cfg := createConnectionConfig()
	_ = cfg.SetKey("group.id", "user-service")
	_ = cfg.SetKey("auto.offset.reset", "earliest")
	producer, err := kafka.NewProducer(createConnectionConfig())
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create producer: %s", err))
	}
	return producer
}

func CreateMessage(data []byte, topic string) *kafka.Message {
	return &kafka.Message{
		Value: data,
		TopicPartition: kafka.TopicPartition{Topic: &topic,
			Partition: kafka.PartitionAny},
	}
}
