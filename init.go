package core

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Initiallizing(c *ConfigStruct, con *ConnectionHandler) {

	logrus.SetLevel(logrus.DebugLevel)

	doneRedis := make(chan struct{}, 1)
	updatesRedis := make(chan *redis.Client, 1)

	doneSQL := make(chan struct{}, 1)
	updatesSQL := make(chan *sqlx.DB, 1)

	doneKafka := make(chan struct{}, 1)
	updatesKafkaProd := make(chan sarama.SyncProducer, 1)
	updatesKafkaCons := make(chan sarama.Consumer, 1)

	doneRabbit := make(chan struct{}, 1)
	updatesRabbit := make(chan *amqp.Connection, 1)

	donePr := make(chan struct{}, 1)
	updatesPr := make(chan *v1.API, 1)
	//sql connection checker
	if c.SqlFlag == "T" {
		go ConToSql(c.SqlHost, c.SqlPort, c.SqlUsername, c.SqlPassword, c.SqlDB, c.SqlDriver, doneSQL, updatesSQL)
		<-doneSQL
		db := <-updatesSQL
		con.SQLDB = db
		con.SQLDBIsInitialized = true
	}
	//kafka connection checker
	if c.KafkaFlag == "T" {
		go ConToKafka(c.KafkaHost, c.KafkaPort, c.KafkaUsername, c.KafkaPassword, doneKafka, updatesKafkaProd, updatesKafkaCons)
		<-doneKafka
		kafkaProd := <-updatesKafkaProd
		kafkaCons := <-updatesKafkaCons
		con.KafkaProducer = kafkaProd
		con.KafkaConsumer = kafkaCons
		con.KafkaIsInitialized = true
	}
	//rabbitMQ connection checker
	if c.RabbitmqFlag == "T" {
		go ConnToRabbitMQ(c.RabbitmqHost, c.RabbitmqPort, c.RabbitmqUsername, c.RabbitmqPassword, doneRabbit, updatesRabbit)
		<-doneRabbit
		conRabbit := <-updatesRabbit
		con.RabbitMQ = conRabbit
		con.RabbitMQIsInitialized = true
	}
	//Prometheus connection checker
	if c.PrometheusFlag == "T" {
		go ConToPrometheus(c.PrometheusHost, c.PrometheusPort, donePr, updatesPr)
		<-donePr
		promCon := <-updatesPr

		con.Prometheus = promCon

	}
	//reddis connection checker
	if c.RedisFlag == "T" {
		go ConToRedis(c.RedisHost, c.RedisPort, c.RedisPassword, doneRedis, updatesRedis)
		<-doneRedis
		client := <-updatesRedis
		con.Redis = client
		con.RedisIsInitialized = true
	}

}
