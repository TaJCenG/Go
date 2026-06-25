package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"user-service/internal/middleware"
	"user-service/internal/user"
)

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.Handle("/users", middleware.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			user.CreateUserHandler(w, r)
		case http.MethodGet:
			user.GetUsersHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	mux.Handle("/users/", middleware.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			user.GetUserHandler(w, r)
		case http.MethodPut:
			user.UpdateUserHandler(w, r)
		case http.MethodDelete:
			user.DeleteUserHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Server config
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Run server in goroutine
	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
