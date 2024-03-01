package core

import (
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"net/http"
	"time"
)

// NewPrometheus creates a new client to interact with the Prometheus API
func NewPrometheus(host, port string) (*v1.API, error) {
	// Forming the Prometheus server address
	address := fmt.Sprintf("http://%s:%s", host, port)

	// Setting up an HTTP client
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Setting up an HTTP client
	clientConfig := api.Config{
		Address:      address,
		RoundTripper: httpClient.Transport,
	}

	// Creating a client for the Prometheus API
	client, err := api.NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Prometheus client: %v", err)
	}

	// Creating API v1 client
	v1api := v1.NewAPI(client)
	return &v1api, nil
}
