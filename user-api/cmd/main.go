package main

import (
	"log"
	"net/http"

	"user-api/internal/account"
	"user-api/internal/config"
	"user-api/internal/middleware"
	"user-api/internal/user"
)

func main() {
	// Step 1: Init DB
	db := config.NewDatabase()
	defer db.Close()

	// Step 2: Create repositories
	userRepo := user.NewRepository(db)
	accountRepo := account.NewRepository(db)

	// Step 3: Create services
	userService := user.NewService(userRepo, db)
	accountService := account.NewService(accountRepo, db)

	// Step 4: Create handlers
	userHandler := user.NewHandler(userService)
	accountHandler := account.NewHandler(accountService)

	// Step 5: Setup routes
	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)
	accountHandler.RegisterRoutes(mux)

	// Step 6: Wrap with middleware
	server := middleware.Recovery(
		middleware.Logging(mux),
	)

	// Step 7: Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
