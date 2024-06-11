package produser

import (
	"fmt"
	"github.com/IvaCheMih/RabbitMQ_pattern/exchange"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func PublishOne(ex exchange.ExchangeStruct, message string, routingKey string) error {

	err := ex.Channel.PublishWithContext(ex.Ctx,
		ex.Name,    // exchange
		routingKey, // routing key
		false,      // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		})
	if err != nil {
		return err
	}

	log.Println(" [x] Sent: ", message)
	fmt.Println()

	return err
}

func PublishMany(ex exchange.ExchangeStruct, messages []string, routingKey string) error {
	var err error

	for _, message := range messages {
		err = PublishOne(ex, message, routingKey)
		if err != nil {
			return err
		}
	}

	return err

}
