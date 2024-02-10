package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func GetReader() *kafka.Consumer {
	cfg := createConnectionConfig()
	_ = cfg.SetKey("group.id", "user-service")
	_ = cfg.SetKey("auto.offset.reset", "earliest")
	c, err := kafka.NewConsumer(cfg)

	if err != nil {
		log.Fatal(err)
	}

	c.SubscribeTopics([]string{"users"}, nil)
	return c
}
