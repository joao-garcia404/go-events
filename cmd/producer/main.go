package main

import "github.com/joao-garcia404/go-events/pkg/rabbitmq"

func main() {
	channel, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer channel.Close()

	rabbitmq.Publish(channel, "message from producer", "amq.direct")
}