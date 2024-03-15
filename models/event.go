package models

import (
	"github.com/amanjshah/event-booking-system/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// Save in-memory til db is introduced
var events = []Event{}

func (e Event) Save() error {
	// Note to self: always execute such queries like this (inject values via Exec method rather than adding values into the query string manually).
	// Protects against SQL injection attacks.
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?,?,?,?,?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() []Event {
	return events
}
