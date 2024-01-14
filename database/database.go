package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var connectionStr string

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionStr = os.Getenv("CONNECTION_STRING")
}

func Init() {
	// Initialize the database connection
	//connStr := "postgresql://Akil-Elias:k7CI6oLEZeAM@ep-calm-leaf-80663857.us-east-1.aws.neon.tech/todoDB?sslmode=require"
	loadEnv()
	var err error
	DB, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}
