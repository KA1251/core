package core

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// NewKafka creates a new connection to Kafka
func NewKafka(host, port, username, password string) (sarama.SyncProducer, sarama.Consumer, error) {
	brokers := []string{host + ":" + port}

	// Setting up Kafka Producer
	producerConfig := sarama.NewConfig()
	producerConfig.Net.SASL.Enable = true
	producerConfig.Net.SASL.User = username
	producerConfig.Net.SASL.Password = password
	// ToDo Additional producerConfig settings

	producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize Kafka Producer: %w", err)
	}

	// Setting up Kafka Consumer
	consumerConfig := sarama.NewConfig()
	consumerConfig.Net.SASL.Enable = true
	consumerConfig.Net.SASL.User = username
	consumerConfig.Net.SASL.Password = password

	// Additional consumerConfig settings

	consumer, err := sarama.NewConsumer(brokers, consumerConfig)
	if err != nil {
		_ = producer.Close() // Close Producer if Consumer failed to initialize
		return nil, nil, fmt.Errorf("failed to initialize Kafka Consumer: %w", err)
	}

	return producer, consumer, nil
}
