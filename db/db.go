package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10) // SQLite does not support concurrent writes, so we limit to 1 connection
	DB.SetMaxIdleConns(5)

	// Create events table if it doesn't exist
	CreateEventsTable()
}

func CreateEventsTable() {
	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,	
		"name" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"location" TEXT NOT NULL,
		"date_time" DATETIME NOT NULL,
		"user_id" INTEGER NOT NULL
	);`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic(err)
	}
}
