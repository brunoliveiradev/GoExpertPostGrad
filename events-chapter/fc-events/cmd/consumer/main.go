package main

import (
	"github.com/brunoliveiradev/GoExpertPostGrad/events-chapter/fc-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"time"
)

func main() {
	ch, err := rabbitmq.SetupChannel()
	if err != nil {
		log.Fatal("Failed to set up channel: ", err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.StartConsumer(ch, "queueName", msgs)

	processAndPublishMessages(ch, msgs)
}

func processAndPublishMessages(ch *amqp.Channel, msgs chan amqp.Delivery) {
	go rabbitmq.ProcessMessages(msgs)

	for i := 0; i < 10; i++ {
		message := generateRandomString(10)
		if err := rabbitmq.Publish(ch, "queueName", message); err != nil {
			log.Printf("Error publishing message: %s", err)
		} else {
			log.Printf("Message published: %s", message)
		}
		time.Sleep(1 * time.Second)
	}
}

func generateRandomString(n int) string {
	rand.NewSource(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
