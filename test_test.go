package main

import (
	"fmt"
	"github.com/IvaCheMih/RabbitMQ_pattern/consumer"
	"github.com/IvaCheMih/RabbitMQ_pattern/exchange"
	"github.com/IvaCheMih/RabbitMQ_pattern/produser"
	"github.com/IvaCheMih/RabbitMQ_pattern/queue"
	"github.com/IvaCheMih/RabbitMQ_pattern/vars"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCase(t *testing.T) {

	var testExchange = vars.ExchangeTest{
		Name:    "exchange_test",
		Kind:    "direct",
		AutoDel: true,
	}

	var testQueue = vars.QueueTest{
		Name:     "queue_test",
		Key:      "ex",
		AutoDel:  true,
		Exchange: testExchange.Name,
	}

	ex, err := exchange.DeclareExchange(testExchange.Name, testExchange.Kind, testExchange.AutoDel, vars.RabbitURL)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Should not produce an error")
	}

	qu, err := queue.CreateQueue(testQueue.Name, testQueue.Key, testQueue.Exchange, testQueue.AutoDel, vars.RabbitURL)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	consumerChan := make(chan vars.MessageInChan)

	go consumer.StartConsumer(qu, "test_consumer", consumerChan)

	err = produser.PublishMany(ex, vars.Messages, testQueue.Key)
	if err != nil {
		t.Errorf("Should not produce an error")
	}

	for i := 0; i < len(vars.Messages); i++ {
		m := <-consumerChan
		if m.Err != nil {
			t.Errorf(m.Err.Error())
			fmt.Println()
		} else {
			log.Println("[*] Message from ", qu.Queue.Name, ": ", string(m.Message))
			fmt.Println()
		}

	}

	assert.Nil(t, err)
	assert.Equal(t, nil, err)

}
