package core

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Initiallizing(c *ConfigStruct, con *ConnectionHandler) {
	//Prometheus
	logrus.SetLevel(logrus.DebugLevel)
	if c.PrometheusFlag == "T" {
		host := c.PrometheusHost
		port := c.PrometheusPort
		client, err := NewPrometheus(host, port)
		if err != nil {
			logrus.Error("Prometheus connection error", err)
		} else {
			logrus.Info("prometheus: connection sucsessfull")
			con.Prometheus = client
			con.PrometheusIsInitialized = true
			Handler.Prometheus = client
			Handler.PrometheusIsInitialized = true
		}
	}
	//Redis
	if c.RedisFlag == "T" {

		addr := c.RedisHost
		port := c.RedisPort
		password := c.RedisPassword
		client, err := NewRedis(addr, port, password)
		if err != nil {
			logrus.Error("Redis connection error:", err)
		} else {
			logrus.Info("redis: connection sucsessfull")
			con.Redis = client
			con.RedisIsInitialized = true
			Handler.Redis = client
			Handler.RedisIsInitialized = true
		}

	}
	//RabbitMQ

	if c.RabbitmqFlag == "T" {
		host := c.RabbitmqHost
		port := c.RabbitmqPort
		user := c.RabbitmqUsername
		password := c.RedisPassword
		conn, err := NewRabbitMQ(host, port, user, password)
		if err != nil {
			logrus.Error("RabbitMQ connection error", err)
		} else {
			logrus.Info("rabbitmq: connection sucsessfull")
			con.RabbitMQ = conn
			con.RabbitMQIsInitialized = true
			Handler.RabbitMQ = conn
			Handler.RabbitMQIsInitialized = true
		}
	}

	//Kafka
	if c.KafkaFlag == "T" {
		host := c.KafkaHost
		port := c.KafkaPort
		username := c.KafkaUsername
		password := c.KafkaPassword

		producer, consumer, err := NewKafka(host, port, username, password)
		if err != nil {
			logrus.Error("Kafka connection error", err)
		} else {
			logrus.Info("kafka: connection sucsessfull")
			con.KafkaConsumer = consumer
			con.KafkaProducer = producer
			con.KafkaIsInitialized = true
			Handler.KafkaProducer = producer
			Handler.KafkaConsumer = consumer
			Handler.KafkaIsInitialized = true
		}

	}
	//SQL
	if c.SqlFlag == "T" {
		host := c.SqlHost
		port := c.SqlPort
		user := c.SqlUsername
		password := c.SqlPassword
		dbName := c.SqlDB
		dbDriver := os.Getenv("SQLDB_DRIVER")
		db, err := NewSqlDB(host, port, user, password, dbName, dbDriver)
		if err != nil {
			logrus.Error("SQL connection error", err)
		} else {
			logrus.Info("SQL: connection sucsessfull")
			con.SQLDB = db
			con.SQLDBIsInitialized = true
			Handler.SQLDB = db
			Handler.SQLDBIsInitialized = true
		}
	}
}
