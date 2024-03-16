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

func (u User) Save() error {
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

func (u User) Validate() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var password string
	err := row.Scan(&password)
	if err != nil {
		return errors.New("Invalid credentials. ")
	}

	// Validate that real password is equivalent to the password sent in the request
	if !utils.CheckPassword(password, u.Password) {
		return errors.New("Invalid credentials. ")
	}
	return nil
}
