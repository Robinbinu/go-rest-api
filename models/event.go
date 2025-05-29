package models

import (
	"errors"
	"log"
	"time"

	"example.com/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,dateTime,user_ID)
	VALUES(?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	//auto close after func
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil

}

func GetAllEvents() ([]Event, error) {
	var events []Event
	query := `
	SELECT * FROM EVENTS
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("no rows found")
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) { //using a pointer to return nil
	var event Event
	query := `
			SELECT * FROM events
			WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, user_ID = ?
	WHERE events.ID = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println("Error during query preparation ", err)
		return err
	}
	//auto close after func
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID,e.ID)
	if err != nil {
		log.Println("Error during query execution ", err)
		return err
	}

	return nil
}
