package core

import "os"

//pushing config to struct
func PushData(c *ConfigStruct) {
	c.RabbitmqFlag = os.Getenv("RABBITMQ_ENABLED")
	c.RabbitmqHost = os.Getenv("RABBITMQ_HOST")
	c.RabbitmqPort = os.Getenv("RABBITMQ_PORT")
	c.RabbitmqUsername = os.Getenv("RABBITMQ_USERNAME")
	c.RabbitmqPassword = os.Getenv("RABBITMQ_PASSWORD")

	c.RedisFlag = os.Getenv("REDIS_ENABLED")
	c.RedisHost = os.Getenv("REDIS_HOST")
	c.RedisPort = os.Getenv("REDIS_PORT")
	c.RedisPassword = os.Getenv("REDDIS_PASSWORD")

	c.KafkaFlag = os.Getenv("KAFKA_ENABLED")
	c.KafkaHost = os.Getenv("KAFKA_HOST")
	c.KafkaPassword = os.Getenv("KAFKA_PASSWORD")
	c.KafkaPort = os.Getenv("KAFKA_PORT")
	c.KafkaUsername = os.Getenv("KAFKA_USERNAME")

	c.SqlFlag = os.Getenv("SQL_ENABLED")
	c.SqlHost = os.Getenv("SQL_HOST")
	c.SqlPassword = os.Getenv("SQL_PASSWORD")
	c.SqlPort = os.Getenv("SQL_PORT")
	c.SqlUsername = os.Getenv("SQL_USERNAME")
	c.SqlDriver = os.Getenv("SQL_DRIVER")
	c.SqlDB = os.Getenv("SQL_DB")

	c.PrometheusFlag = os.Getenv("PROMETHEUS_ENABLED")
	c.PrometheusHost = os.Getenv("PROMETHEUS_HOST")
	c.PrometheusPort = os.Getenv("PROMETHEUS_PORT")

}
