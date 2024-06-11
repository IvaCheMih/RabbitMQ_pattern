package produser

import (
	"fmt"
	"github.com/IvaCheMih/RabbitMQ_pattern/exchange"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func PublishOne(ex exchange.ExchangeStruct, message string, routingKey string) {

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
		log.Println(err)
	}

	log.Println(" [x] Sent: ", message)
	fmt.Println()

}

func PublishMany(ex exchange.ExchangeStruct, messages []string, routingKey string) {

	for _, message := range messages {
		PublishOne(ex, message, routingKey)
	}

}
