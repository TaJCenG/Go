package health

import (
	"net/http"
)

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("alive"))
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	// In production, check DB, Redis, Kafka connections
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
}
