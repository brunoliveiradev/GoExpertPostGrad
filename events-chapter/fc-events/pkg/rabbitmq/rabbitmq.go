package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func SetupChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Println("Failed to open a channel:", err)
		return nil, err
	}
	return ch, nil
}

func StartConsumer(ch *amqp.Channel, queueName string, msgs chan<- amqp.Delivery) {
	if err := ensureQueueExists(ch, queueName); err != nil {
		log.Fatalf("Failed to ensure queue exists: %s", err)
		return
	}

	if err := consumeMessages(ch, queueName, msgs); err != nil {
		log.Fatalf("Failed to start consumer: %s", err)
	}
}

func ProcessMessages(msgs chan amqp.Delivery) {
	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		if err := msg.Ack(false); err != nil {
			log.Printf("Failed to acknowledge message: %s", err)
		}
		if string(msg.Body) == "exit" {
			log.Println("Exit message received, shutting down...")
			break
		}
	}
}

func Publish(ch *amqp.Channel, queueName string, message string) error {
	if err := ch.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	); err != nil {
		log.Printf("Failed to publish message: %s", err)
		return err
	}
	log.Printf("Message published: %s", message)
	return nil
}

func ensureQueueExists(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare queue: %s", err)
	}
	return err
}

func consumeMessages(ch *amqp.Channel, queueName string, output chan<- amqp.Delivery) error {
	messages, err := ch.Consume(
		queueName,
		"consumer-"+queueName, // Unique consumer tag
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

	log.Printf("Consumer set up successfully for queue: %s", queueName)
	go func() {
		for msg := range messages {
			output <- msg
		}
	}()

	return nil
}
