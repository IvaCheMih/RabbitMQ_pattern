package queue

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Connect *amqp091.Connection
	Channel *amqp091.Channel
	Ctx     context.Context
	Queue   amqp091.Queue
}
