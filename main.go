package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=6465 dbname=Users host=localhost port=5432 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	connStr := "user=postgres dbname=Users password=6465 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")
	defer db.Close()

}

func createTable(db *sql.DB) {
	// Create the "people" table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS people (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			age INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}
//comment
