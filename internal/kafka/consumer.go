package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/third-place/image-service/internal/db"
	"github.com/third-place/image-service/internal/mapper"
	"github.com/third-place/image-service/internal/model"
	"github.com/third-place/image-service/internal/repository"
	"github.com/google/uuid"
	"log"
)

func InitializeAndRunLoop() {
	reader := GetReader()
	userRepository := repository.CreateUserRepository(db.CreateDefaultConnection())
	err := loopKafkaReader(userRepository, reader)
	if err != nil {
		log.Fatal("exited kafka loop due to error :: ", err)
	}
}

func loopKafkaReader(userRepository *repository.UserRepository, reader *kafka.Consumer) error {
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
		userEntity, err := userRepository.FindOneByUuid(uuid.MustParse(userModel.Uuid))
		if err != nil {
			userEntity = mapper.GetUserEntityFromModel(userModel)
			userRepository.Create(userEntity)
		}
	}
}
