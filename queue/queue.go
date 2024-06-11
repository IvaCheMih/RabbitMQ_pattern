package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateQueue(name string, key string, exchange string, autoDel bool, rabbitURL string) (Queue, error) {
	var qu Queue
	var err error

	qu.Connect, err = amqp.Dial(rabbitURL)
	if err != nil {
		return Queue{}, err
	}

	qu.Channel, err = qu.Connect.Channel()
	if err != nil {
		return Queue{}, err
	}

	qu.Queue, err = qu.Channel.QueueDeclare(
		name,    // name
		true,    // durable
		autoDel, // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return Queue{}, err
	}

	err = qu.Channel.QueueBind(name, key, exchange, false, nil)
	if err != nil {
		return Queue{}, err
	}

	err = qu.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	return qu, err
}
