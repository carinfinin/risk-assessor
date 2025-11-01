package mq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/carinfinin/risk-assessor/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	MQ      *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func New(cfg *config.Config) (*Producer, error) {
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
		_ = conn.Close()
		_ = channel.Close()
		return nil, err
	}
	return &Producer{
		MQ:      conn,
		Channel: channel,
		Queue:   q,
	}, err
}

func (r *Producer) Stop() error {
	err := r.Channel.Close()
	if err != nil {
		return err
	}
	err = r.MQ.Close()
	return err
}

func (r *Producer) Send(ctx context.Context, data *model.ClientData) error {
	ctxTime, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.Channel.PublishWithContext(ctxTime,
		"",
		r.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        dataJson,
		},
	)
}
