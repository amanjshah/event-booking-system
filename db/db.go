package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database. ")
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
	    "event_name" TEXT NOT NULL,
		"description" TEXT NOT NULL,
	 	"location" TEXT NOT NULL,
	    "dataTime" DATETIME NOT NULL,
	    "user_id" INTEGER
	);`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table. ")
	}

}
