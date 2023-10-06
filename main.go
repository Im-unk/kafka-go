package main

import (
	"fmt"
	"kafka/api"
	"kafka/internal/consumer"
	"kafka/internal/producer"
	"kafka/service"
	"log"
	"net/http"
)

func main() {
	brokerAddress := "localhost:9092"
	topic := "my-topic"
	groupID := "my-consumer-group2"

	producer, err := producer.NewProducer(brokerAddress)
	if err != nil {
		fmt.Printf("Error Creating a Producer %v \n ", err)
		return
	}

	consumer, err := consumer.NewConsumer(brokerAddress, topic, groupID)
	if err != nil {
		fmt.Printf("Error Creating a Producer %v \n ", err)
		return
	}

	service := service.NewService(producer, consumer)

	go consumer.Consume()

	router := api.NewRouter(service)
	log.Fatal(http.ListenAndServe(":8080", router))

}
