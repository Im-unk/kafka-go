package service

import (
	"fmt"
	"kafka/internal/consumer"
	"kafka/internal/producer"
	"kafka/models"
)

type Service struct {
	producer *producer.Producer
	consumer *consumer.Consumer
}

func NewService(producer *producer.Producer, consumer *consumer.Consumer) *Service {
	return &Service{
		producer: producer,
		consumer: consumer,
	}
}

func Get() error {
	return nil
}

func (s *Service) Create(data models.Data) (models.Data, error) {

	if data.Partition != 0 {
		producedData := s.producer.ProduceWithPartition(data.Topic, int32(data.Partition), data.Key, data.Value)
		fmt.Println(producedData)

	} else {
		producedData := s.producer.Produce(data.Topic, data.Key, data.Value)
		fmt.Println(producedData)

	}

	return data, nil
}
