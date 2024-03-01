package core

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

// StartHealthCheckServer starts the HTTP server for health check
func StartHealthCheckServer(port string) {
	logrus.SetLevel(logrus.DebugLevel)
	http.HandleFunc("/health", healthCheckHandler)
	log.Println("Starting Health Check Server on port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logrus.Error("Health Check Server failed", err)
	}
}

// healthCheckHandler processes status check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the service status
	//add separation

	// If all checks are successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
