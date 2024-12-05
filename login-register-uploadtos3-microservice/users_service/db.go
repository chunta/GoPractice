package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("InitDB")
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(6)

	if err := createTables(); err != nil {
		panic("Could not create tables: " + err.Error())
	}
}

func createTables() error {
	// Define the SQL for creating the Users table
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        hashed_password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
    );`

	// Execute the SQL command
	_, err := DB.Exec(createUsersTable)
	return err
}
