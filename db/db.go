package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var Database *sql.DB

func InitDatabase() {
	var err error
	Database, err = sql.Open("sqlite", "api.db")
	if err != nil {
		fmt.Println(err)
		panic("could not connect to the database.")
	}

	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		dateTime datetime NOT NULL,
		user_id INTEGER
	)
	`

	_, err := Database.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("could not create the events table.")
	}
}
