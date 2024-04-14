package main

import (
	"github.com/brunoliveiradev/GoExpertPostGrad/events-chapter/fc-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	ch, err := setupChannel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, msgs)

	for msg := range msgs {
		processMessage(msg)
	}
}

func setupChannel() (*amqp.Channel, error) {
	return rabbitmq.OpenChannel()
}

func processMessage(msg amqp.Delivery) {
	log.Printf("Received a message: %s", msg.Body)
	msg.Ack(false)

	if string(msg.Body) == "exit" {
		log.Fatal("Exit message received")
	}
}
