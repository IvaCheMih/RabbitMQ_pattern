package consumer

import (
	"fmt"
	"github.com/IvaCheMih/RabbitMQ_pattern/queue"
	"github.com/IvaCheMih/RabbitMQ_pattern/vars"
	"log"
)

func StartConsumer(queue queue.Queue, consumerName string, consumerChan chan vars.MessageInChan) {

	msgs, err := queue.Channel.Consume(
		queue.Queue.Name, // queue
		"",               // consumer
		false,            // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		consumerChan <- vars.MessageInChan{[]byte{}, err}
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {

			log.Printf("Received a message: %s", d.Body)
			log.Println("Done", consumerName, " from queue: ", queue.Queue.Name)
			fmt.Println()
			er := d.Ack(true)

			consumerChan <- vars.MessageInChan{d.Body, er}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
