package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type UserDTO struct {
	USER_ID    int
	FIRST_NAME string
	LAST_NAME  string
	EMAIL      string
}

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	username := os.Getenv(DB_USERNAME)
	password := os.Getenv(DB_PASSWORD)

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, username, password, DB_NAME)

	db, err := sqlx.Connect("postgres", psqlConn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	getUsers(db)

	fmt.Println("Connected to the database!")
}

func getUsers(db *sqlx.DB) []UserDTO {
	sqlQuery := "SELECT * FROM USERS"

	users := []UserDTO{}
	err := db.Select(&users, sqlQuery)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
