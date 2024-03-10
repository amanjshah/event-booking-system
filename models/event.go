package models

import "time"

type Event struct {
	ID          int    `binding:"required"`
	Name        string `binding:"required"`
	Description string
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

// Save in-memory til db is introduced
var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
