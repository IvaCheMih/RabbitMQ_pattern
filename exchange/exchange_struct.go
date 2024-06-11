package exchange

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

type ExchangeStruct struct {
	Name    string
	Connect *amqp091.Connection
	Channel *amqp091.Channel
	Ctx     context.Context
}
