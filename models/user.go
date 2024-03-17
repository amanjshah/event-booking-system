package models

import (
	"errors"
	"github.com/amanjshah/event-booking-system/db"
	"github.com/amanjshah/event-booking-system/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var id int64
	var password string
	err := row.Scan(&id, &password)
	if err != nil {
		return errors.New("Invalid credentials. ")
	}

	u.ID = id
	// Validate that real password is equivalent to the password sent in the request
	if !utils.CheckPassword(password, u.Password) {
		return errors.New("Invalid credentials. ")
	}
	return nil
}

func (u *User) getEmailFromId() (string, error) {
	query := "SELECT email FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, u.ID)

	var email string
	err := row.Scan(&email)
	if err != nil {
		return "", errors.New("Invalid email. ")
	}

	return email, nil
}
