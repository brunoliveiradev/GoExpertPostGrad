package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type RabbitMQConfig struct {
	URL      string
	Queue    string
	Exchange string
}

func SetupChannel(config RabbitMQConfig) (*amqp.Channel, *amqp.Connection, error) {
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Printf("Failed to open a channel: %v", err)
		return nil, nil, err
	}
	return ch, conn, nil
}

func StartConsumer(ch *amqp.Channel, config RabbitMQConfig, msgs chan<- amqp.Delivery) error {
	if err := ensureQueueExists(ch, config.Queue); err != nil {
		return fmt.Errorf("failed to ensure queue exists: %w", err)
	}

	if err := consumeMessages(ch, config.Queue, msgs); err != nil {
		return fmt.Errorf("failed to start consumer: %w", err)
	}
	return nil
}

func ProcessMessages(msgs chan amqp.Delivery) {
	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		if err := msg.Ack(false); err != nil {
			log.Printf("Failed to acknowledge message: %v", err)
		}
		if string(msg.Body) == "exit" {
			log.Println("Exit message received, shutting down...")
			break
		}
	}
}

func Publish(ch *amqp.Channel, config RabbitMQConfig, message string) error {
	if err := ch.PublishWithContext(
		context.Background(),
		config.Exchange,
		config.Queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	log.Printf("Message published: %s", message)
	return nil
}

func ensureQueueExists(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	return nil
}

func consumeMessages(ch *amqp.Channel, queueName string, output chan<- amqp.Delivery) error {
	messages, err := ch.Consume(
		queueName,
		"consumer-"+queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error creating consumer with RabbitMQ: %w", err)
	}

	log.Printf("Consumer set up successfully for queue: %s", queueName)
	go func() {
		for msg := range messages {
			output <- msg
		}
	}()

	return nil
}

func SetupGracefulShutdown(conn *amqp.Connection, ch *amqp.Channel) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ch.Close()
		conn.Close()
		log.Println("RabbitMQ connection closed gracefully")
		os.Exit(0)
	}()
}
