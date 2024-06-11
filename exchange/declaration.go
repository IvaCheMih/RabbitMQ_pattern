package exchange

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func DeclareExchange(name string, kind string, autoDel bool, rabbitURL string) (ExchangeStruct, error) {
	var ex ExchangeStruct
	var err error

	ex.Name = name

	ex.Connect, err = amqp.Dial(rabbitURL) // Создаем подключение к RabbitMQ
	if err != nil {
		return ExchangeStruct{}, err
	}

	ex.Channel, err = ex.Connect.Channel()
	if err != nil {
		return ExchangeStruct{}, err
	}

	err = ex.Channel.ExchangeDeclare(name, kind, true, autoDel, false, false, nil)
	if err != nil {
		return ExchangeStruct{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ex.Ctx = ctx

	return ex, err
}
