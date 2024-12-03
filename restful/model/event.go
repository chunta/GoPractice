package model

import (
	"fmt"
	"restful/db"
	"time"
)

type Event struct {
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
    Des       string    `json:"description"`
    CreatedAt time.Time `json:"created_at"`
}


func (e Event) Save() error {
	fmt.Println("Save...")

	query := `
	INSERT INTO events(name, description) VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Des)
	if err != nil {
		fmt.Println(err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	fmt.Println("Get All Events")

	var events []Event

	// Define the SQL query
	query := `SELECT * FROM events`

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	
	fmt.Println(rows)

	// Iterate through the rows and scan the data into the events slice
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Des, &event.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		events = append(events, event)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return events, nil
}
