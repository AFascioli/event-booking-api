package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDatabase() {
	var err error
	Database, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not connect to database.")
	}

	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := Database.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = Database.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = Database.Exec(createRegistrationsTable)
	if err != nil {
		panic("could not create registrations table")
	}

}
