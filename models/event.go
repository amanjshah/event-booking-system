package models

import (
	"github.com/amanjshah/event-booking-system/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// Note to self: always execute such queries like this (inject values via Exec method rather than adding values into the query string manually).
	// Protects against SQL injection attacks.
	query := `
	INSERT INTO events(event_name, description, location, user_id) 
	VALUES(?,?,?,?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.UserID)
	if err != nil {
		return err
	}
	// Use LastInsertID() method to get the autogenerated ID
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// Use Query when you want to query data, and Exec when you want to insert/update data.
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event
	// loop until there are no more rows
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET event_name = ?, description = ?, location = ?
	WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.Name, e.Description, e.Location, e.ID)
	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}
