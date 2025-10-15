package models

import (
	"errors"

	"github.com/AqeelMohammed-programmer/event-booking-api/db"
	"github.com/AqeelMohammed-programmer/event-booking-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"

	reslut := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := reslut.Scan(&retrivedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isValidPassword := utils.CompareHashedPassword(u.Password, retrivedPassword)

	if !isValidPassword {
		return errors.New("invalid credentials")
	}

	return nil
}
