package core

import (
	"os"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Initiallizing(c *ConfigStruct, con *ConnectionHandler) {

	logrus.SetLevel(logrus.DebugLevel)

	addrRedis := c.RedisHost
	portRedis := c.RedisPort
	passwordRedis := c.RedisPassword
	doneRedis := make(chan struct{}, 1)
	updatesRedis := make(chan *redis.Client, 1)

	hostSQL := c.SqlHost
	portSQL := c.SqlPort
	userSQL := c.SqlUsername
	passwordSQL := c.SqlPassword
	dbNameSQL := c.SqlDB
	dbDriverSQL := os.Getenv("SQLDB_DRIVER")
	doneSQL := make(chan struct{}, 1)
	updatesSQL := make(chan *sqlx.DB, 1)

	hostKafka := c.KafkaHost
	portKafka := c.KafkaPort
	usernameKafka := c.KafkaUsername
	passwordKafka := c.KafkaPassword
	doneKafka := make(chan struct{}, 1)
	updatesKafkaProd := make(chan sarama.SyncProducer, 1)
	updatesKafkaCons := make(chan sarama.Consumer, 1)

	hostRabbit := c.RabbitmqHost
	portRabbit := c.RabbitmqPort
	userRabbit := c.RabbitmqUsername
	passwordRabbit := c.RedisPassword
	doneRabbit := make(chan struct{}, 1)
	updatesRabbit := make(chan *amqp.Connection, 1)

	hostPr := c.PrometheusHost
	portPr := c.PrometheusPort
	donePr := make(chan struct{}, 1)
	updatesPr := make(chan *v1.API, 1)
	//sql connection checker
	if c.SqlFlag == "T" {
		go ConToSql(hostSQL, portSQL, userSQL, passwordSQL, dbNameSQL, dbDriverSQL, doneSQL, updatesSQL)
		<-doneSQL
		db := <-updatesSQL
		con.SQLDB = db
		con.SQLDBIsInitialized = true
	}
	//kafka connection checker
	if c.KafkaFlag == "T" {
		go ConToKafka(hostKafka, portKafka, usernameKafka, passwordKafka, doneKafka, updatesKafkaProd, updatesKafkaCons)
		<-doneKafka
		kafkaProd := <-updatesKafkaProd
		kafkaCons := <-updatesKafkaCons
		con.KafkaProducer = kafkaProd
		con.KafkaConsumer = kafkaCons
		con.KafkaIsInitialized = true
	}
	//rabbitMQ connection checker
	if c.RabbitmqFlag == "T" {
		go ConnToRabbitMQ(hostRabbit, portRabbit, userRabbit, passwordRabbit, doneRabbit, updatesRabbit)
		<-doneRabbit
		conRabbit := <-updatesRabbit
		con.RabbitMQ = conRabbit
		con.RabbitMQIsInitialized = true
	}
	//Prometheus connection checker
	if c.PrometheusFlag == "T" {
		go ConToPrometheus(hostPr, portPr, donePr, updatesPr)
		<-donePr
		promCon := <-updatesPr

		con.Prometheus = promCon

	}
	//reddis connection checker
	if c.RedisFlag == "T" {
		go ConToRedis(addrRedis, portRedis, passwordRedis, doneRedis, updatesRedis)
		<-doneRedis
		client := <-updatesRedis
		con.Redis = client
		con.RedisIsInitialized = true
	}

}
