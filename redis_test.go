package core

import (
	"os"
	"testing"
)

func TestRedisInitialization(t *testing.T) {
	// ToDo add a server for tests and set credentials in the code below
	//// Preserving the original values of environment variables
	//originalHost := os.Getenv("REDIS_HOST")
	//originalPort := os.Getenv("REDIS_PORT")
	//originalPassword := os.Getenv("REDIS_PASSWORD")
	//
	//// Setting test values of environment variables
	//os.Setenv("REDIS_HOST", "test_host")
	//os.Setenv("REDIS_PORT", "test_port")
	//os.Setenv("REDIS_PASSWORD", "test_password")
	//
	//defer func() {
	//	// Restoring the original values of environment variables after the test
	//	os.Setenv("REDIS_HOST", originalHost)
	//	os.Setenv("REDIS_PORT", originalPort)
	//	os.Setenv("REDIS_PASSWORD", originalPassword)
	//}()

	// Getting the values of environment variables for initialization
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	// Initializing the Redis client with parameters
	redisClient, err := NewRedis(host, port, password)
	if err != nil {
		t.Errorf("Error initializing Redis: %v", err)
	}

	// Verifying that the Redis client was successfully created
	if redisClient == nil {
		t.Error("Redis client not created")
	}
}
