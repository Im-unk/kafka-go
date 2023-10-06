package consumer

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	kafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	Consumer *kafka.Consumer
}

func NewConsumer(brokerAddress, topic, groupID string) (*Consumer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": brokerAddress,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		Consumer: c,
	}, nil
}

func (c *Consumer) Consume() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	run := true
	for run {
		// select {
		// case ev := <-c.Consumer.Events():
		// 	switch e := ev.(type) {
		// 	case *kafka.Message:
		// 		fmt.Printf("Consumed message from topic %s [%d] at offset %v: %s \n", *e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Value))
		// 	case kafka.Error:
		// 		fmt.Printf("Error in Consumption %v: \n", e)
		// 	}
		// case <-sigchan:
		// 	fmt.Println("Consumer Interrupted, Shutting Down ...")
		// 	run = false
		// 	c.Consumer.Close()
		// }

		msg, err := c.Consumer.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error %v (%v) \n", err, msg)
		}
	}

	c.Consumer.Close()
}
