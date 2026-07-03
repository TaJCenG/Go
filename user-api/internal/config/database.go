package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDatabase() *sql.DB {
	godotenv.Load()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Connection pool tuning
	db.SetMaxOpenConns(25)                 // max concurrent connections
	db.SetMaxIdleConns(25)                 // keep idle connections ready
	db.SetConnMaxLifetime(5 * time.Minute) // recycle connections after 5 min
	db.SetConnMaxIdleTime(1 * time.Minute) // close idle after 1 min

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}
