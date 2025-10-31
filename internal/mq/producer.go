package mq

import (
	"github.com/carinfinin/risk-assessor/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	MQ     *amqp.Connection
	Chanel *amqp.Channel
	Queue  amqp.Queue
}

func New() *Producer {
	return &Producer{}
}

func (r *Producer) Inicialize(cfg *config.Config) error {
	conn, err := amqp.Dial(cfg.MQPath)
	if err != nil {
		return err
	}

	chanel, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := chanel.QueueDeclare(
		"event",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	r.MQ = conn
	r.Chanel = chanel
	r.Queue = q

	return nil
}

func (r *Producer) Stop() error {
	err := r.MQ.Close()
	if err != nil {
		return err
	}
	err = r.Chanel.Close()
	return err
}
