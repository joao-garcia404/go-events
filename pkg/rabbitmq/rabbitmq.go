package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel();

	if err != nil {
		panic(err)
	}

	return channel, nil;
}

func Consume(channel *amqp.Channel, out chan amqp.Delivery) error {
	messages, err := channel.Consume(
		"queue",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for message := range messages {
		out <- message
	}

	return nil
}