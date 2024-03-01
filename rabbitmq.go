package core

import (
	"fmt"
	"github.com/streadway/amqp"
)

// NewRabbitMQ creates a new connection to RabbitMQ
func NewRabbitMQ(host, port, user, password string) (*amqp.Connection, error) {
	// Generating a connection string for RabbitMQ
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	// Establishing a connection
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	return conn, nil
}
