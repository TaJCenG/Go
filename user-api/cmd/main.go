package main

import (
	"Day1Utils/user-api/internal/account"
	"Day1Utils/user-api/internal/config"
	"Day1Utils/user-api/internal/kafka"
	"Day1Utils/user-api/internal/middleware"
	"Day1Utils/user-api/internal/user"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := config.NewLogger()
	defer logger.Sync()

	logger.Info("Application startup")
	// Step 1: Init DB
	db := config.NewDatabase()
	defer db.Close()
	logger.Info("Database connection established")
	rdb := config.NewRedisClient()

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, db, logger)
	sessionStore := user.NewSessionStore(rdb)

	// Step 2: Wire dependencies
	userRepo := user.NewRepository(db)
	accountRepo := account.NewRepository(db)

	userService := user.NewService(userRepo)
	accountService := account.NewService(accountRepo, db)

	userHandler := user.NewHandler(userService)
	accountHandler := account.NewHandler(accountService)

	// Step 3: Routes
	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)
	accountHandler.RegisterRoutes(mux)

	// Step 4: Middleware
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Recovery(middleware.Logging(mux)),
	}

	// Step 5: Run server in goroutine
	go func() {
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	ctx := context.Background()
	consumer := kafka.NewConsumer([]string{"localhost:9092"}, "orders", "inventory-service")

	consumer.Start(ctx, func(msg string) error {
		// Deduct stock based on order event
		log.Printf("Processing order event: %s", msg)
		return nil
	})
	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)
	accountHandler.RegisterRoutes(mux)

	// Health endpoints
	mux.HandleFunc("/health/live", health.Liveness)
	mux.HandleFunc("/health/ready", health.Readiness)

	// Step 6: Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
