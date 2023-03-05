package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	username := os.Getenv(DB_USERNAME)
	password := os.Getenv(DB_PASSWORD)

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, username, password, DB_NAME)

	db, err := sql.Open("postgres", psqlConn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database!")
}
