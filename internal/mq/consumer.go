package mq

import (
	"encoding/json"
	"fmt"

	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/carinfinin/risk-assessor/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	MQ      *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewConsumer(cfg *config.Config) (*Consumer, error) {
	conn, err := amqp.Dial(cfg.MQPath)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	q, err := channel.QueueDeclare(
		"event",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = channel.Close()
		_ = conn.Close()
		return nil, err
	}

	consumer := &Consumer{
		MQ:      conn,
		Channel: channel,
		Queue:   q,
	}

	return consumer, nil
}
func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	return c.Channel.Consume(
		c.Queue.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}
func (c *Consumer) ProcessMessage(delivery amqp.Delivery) (*model.ClientData, error) {
	// Проверяем content-type
	if delivery.ContentType != "application/json" {
		return nil, fmt.Errorf("unexpected content-type: %s", delivery.ContentType)
	}

	// Декодируем JSON
	var clientData model.ClientData
	if err := json.Unmarshal(delivery.Body, &clientData); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &clientData, nil
}

func (c *Consumer) Stop() error {
	if c.Channel != nil {
		c.Channel.Close()
	}
	if c.MQ != nil {
		return c.MQ.Close()
	}
	return nil
}
