package database

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql" // MySQL driver
// )

// var DB *sql.DB

// // InitDB initializes the database connection
// func InitDB() {
// 	// Read environment variables (production best practice)
// 	// user := os.Getenv("DB_USER")
// 	// pass := os.Getenv("DB_PASS")
// 	// host := os.Getenv("DB_HOST")
// 	// port := os.Getenv("DB_PORT")
// 	// name := os.Getenv("DB_NAME")

// 	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
// 	// 	user, pass, host, port, name)
// 	dsn := "root:1234@tcp(localhost:3306)/user_service?parseTime=true"

// 	var err error
// 	log.Printf("Connecting with DSN: %s", dsn)

// 	DB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatalf("Error opening database: %v", err)
// 	}

// 	// Connection pool settings
// 	DB.SetMaxOpenConns(25)
// 	DB.SetMaxIdleConns(25)
// 	DB.SetConnMaxLifetime(5 * 60) // 5 minutes

// 	// Test connection
// 	if err := DB.Ping(); err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}

// 	log.Println("Database connection established")
// }
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	if !loadEnvFile() {
		log.Println("No .env file found, using system environment variables")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	if missing := missingDBEnvVars(); len(missing) > 0 {
		log.Fatalf("Missing required database environment variables: %s", strings.Join(missing, ", "))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, name)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connection established")
}

func loadEnvFile() bool {
	for _, path := range []string{".env", "../.env"} {
		if err := godotenv.Load(path); err == nil {
			return true
		}
	}
	return false
}

func missingDBEnvVars() []string {
	required := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}
	var missing []string
	for _, key := range required {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}
	return missing
}
