package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

// Save in-memory til db is introduced
var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}
