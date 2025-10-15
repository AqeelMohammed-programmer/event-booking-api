package models

import (
	"database/sql"
	"time"

	"github.com/AqeelMohammed-programmer/event-booking-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

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

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var events []Event

	for result.Next() {
		var event Event

		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description,
		&event.Location, &event.DateTime, &event.UserId); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ? 
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (u *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.ID, userId)

	return err
}
