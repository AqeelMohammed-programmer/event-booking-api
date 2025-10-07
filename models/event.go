package models

import (
	"time"

	"github.com/AqeelMohammed-programmer/event-booking-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events []Event = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	e.ID, err = result.LastInsertId()
	return err
}

func GetAllEvents() []Event {
	return events
}
