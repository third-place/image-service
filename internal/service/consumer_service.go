package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/third-place/image-service/internal/kafka"
	"github.com/third-place/image-service/internal/model"
	"log"
)

type ConsumerService struct {
	userService *UserService
}

func CreateConsumerService() *ConsumerService {
	return &ConsumerService{
		CreateUserService(),
	}
}

func (c *ConsumerService) InitializeAndRunLoop() {
	reader := kafka2.GetReader()
	err := c.loopKafkaReader(reader)
	if err != nil {
		log.Fatal("exited kafka loop due to error :: ", err)
	}
}

func (c *ConsumerService) loopKafkaReader(reader *kafka.Consumer) error {
	for {
		log.Print("listening for kafka messages")
		data, err := reader.ReadMessage(-1)
		if err != nil {
			log.Print("error reading kafka messages :: ", err)
			return nil
		}
		log.Print("consuming user message ", string(data.Value))
		userModel, err := model.DecodeMessageToUser(data.Value)
		if err != nil {
			log.Print("error decoding message to user, skipping", string(data.Value))
			continue
		}
		c.userService.UpsertUser(userModel)
	}
}
