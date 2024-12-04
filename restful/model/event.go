package model

import (
	"database/sql"
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

	// Define the SQL query
	query := `SELECT * FROM events`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	
	fmt.Println(rows)

	// Iterate through the rows and scan the data into the events slice
	var events []Event
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

func GetEventByID(id int64) (*Event, error) {
    fmt.Println("Get Event By ID")

    var event Event

    // Define the SQL query
    query := `SELECT id, name, description, created_at FROM events WHERE id = ?`

    // Execute the query
    row := db.DB.QueryRow(query, id)

    // Scan the result into the Event struct
    err := row.Scan(&event.ID, &event.Name, &event.Des, &event.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            // Return a custom error message if no rows were found
            return nil, fmt.Errorf("event with ID %d not found", id)
        }
        // Return any other errors encountered during the scan
        return nil, err
    }

    return &event, nil
}

func UpdateEvent(event Event) error {
	query := `
	UPDATE events 
	SET name = ?, description = ? 
	WHERE id = ?
	`

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(event.Name, event.Des, event.ID)
	return err
}

func DeleteEventByID(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}