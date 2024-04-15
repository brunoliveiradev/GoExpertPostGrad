package main

import (
	"github.com/brunoliveiradev/GoExpertPostGrad/events-chapter/fc-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	config := rabbitmq.Config{
		URL:      getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		Queue:    getEnv("RABBITMQ_QUEUE", "queueName"),
		Exchange: getEnv("RABBITMQ_EXCHANGE", ""), // Exchange name or leave as empty string for default
	}

	ch, conn, err := rabbitmq.SetupChannel(config)
	if err != nil {
		log.Fatalf("Failed to set up channel: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	rabbitmq.SetupGracefulShutdown(conn, ch)

	msgs := make(chan amqp.Delivery)
	if err := rabbitmq.StartConsumer(ch, config, msgs); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	processAndPublishMessages(ch, config, msgs)
}

func processAndPublishMessages(ch *amqp.Channel, config rabbitmq.Config, msgs chan amqp.Delivery) {
	go rabbitmq.ProcessMessages(msgs)

	for i := 0; i < 10; i++ {
		message := generateRandomString(10)
		if err := rabbitmq.Publish(ch, config, message); err != nil {
			log.Printf("Error publishing message: %v", err)
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

// getEnv retrieves environment variables or returns a default value if not set
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
