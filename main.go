package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.Wrap(err, "environment variables")
	}

	amqpConn, err := newRabbitMQConn()
	if err != nil {
		return errors.Wrap(err, "rabbitmq")
	}
	defer amqpConn.Close()

	srv := newServer(amqpConn)
	return srv.run()
}

// The server type contains the dependencies of our server.
type server struct {
	amqpConn *amqp.Connection
}

func newServer(amqpConn *amqp.Connection) *server {
	return &server{amqpConn: amqpConn}
}

func (s *server) run() error {
	return nil
}
