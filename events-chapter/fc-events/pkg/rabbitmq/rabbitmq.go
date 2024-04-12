package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Error creating connection with RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error creating channel with RabbitMQ:", err)
	}
	return ch, nil
}

func Consumer(ch *amqp.Channel, output chan<- amqp.Delivery) error {
	messages, err := ch.Consume(
		"queueName",
		"consumer-queueName",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Error creating consumer with RabbitMQ:", err)
		return err
	}

	for msg := range messages {
		output <- msg
	}

	return nil
}
