package main

import (
	"log"

	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/carinfinin/risk-assessor/internal/mq"
)

func main() {
	cfg := config.New("")

	consumer, err := mq.NewConsumer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Stop()

	messages, err := consumer.Consume()
	if err != nil {
		log.Fatal(err)
	}

	for delivery := range messages {
		clientData, err := consumer.ProcessMessage(delivery)
		if err != nil {
			log.Printf("Error processing message: %v", err)
			continue
		}

		log.Printf("Received client data: %+v", clientData)
		// Обрабатываем данные...
	}
}
