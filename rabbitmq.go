package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

// Initialize new RabbitMQ connection
func newRabbitMQConn() (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)
	return amqp.Dial(connAddr)
}
