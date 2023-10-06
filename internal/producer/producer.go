package producer

import (
	"fmt"

	kafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	Producer *kafka.Producer
}

func NewProducer(brokerAddress string) (*Producer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": brokerAddress,
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, err
	}

	return &Producer{
		Producer: producer,
	}, nil

}

func (kp *Producer) Produce(topic string, key string, value string) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(key),
		Value: []byte(value),
	}

	err := kp.Producer.Produce(msg, nil)
	if err != nil {
		return err
	}

	deliveryReport := <-kp.Producer.Events()
	m := deliveryReport.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	} else {
		fmt.Printf("Message Produced to topic: %s [%d] with the offset: %v \n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	return nil

}

func (kp *Producer) ProduceWithPartition(topic string, partition int32, key string, value string) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: partition,
		},
		Key:   []byte(key),
		Value: []byte(value),
	}

	err := kp.Producer.Produce(msg, nil)
	if err != nil {
		return err
	}

	deliveryReport := <-kp.Producer.Events()
	m := deliveryReport.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	} else {
		fmt.Printf("Message Produced to topic: %s [%d] with the offset: %v \n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	return nil

}

func (kp *Producer) Close() {
	kp.Producer.Close()
}
