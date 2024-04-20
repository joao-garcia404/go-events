package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

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

func Consume(channel *amqp.Channel, out chan amqp.Delivery, queue string) error {
	messages, err := channel.Consume(
		queue,
		"go-consumer",
		false, // message auto-commit
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

func Publish(channel *amqp.Channel, body string, exchange string) error {
	ctx := context.Background()

	err := channel.PublishWithContext(
		ctx,
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
