package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	// Initialize the database connection
	connStr := "postgresql://Akil-Elias:k7CI6oLEZeAM@ep-calm-leaf-80663857.us-east-1.aws.neon.tech/todoDB?sslmode=require"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}
