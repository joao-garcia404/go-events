package main

import (
	"fmt"

	"github.com/joao-garcia404/go-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	channel, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer channel.Close()

	messages := make(chan amqp.Delivery)

	go rabbitmq.Consume(channel, messages)

	for msg := range messages {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
