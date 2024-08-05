package db

import (
	"database/sql"
	"fmt"
	"os"
)

func DbConnect() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL+"?sslmode=disable")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return nil, err
	}
	fmt.Println("Connected to database successfully!")
	return db, nil
}
