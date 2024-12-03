package note

import (
	"encoding/json"
	"os"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

const fileName = "tmp.json"

func (note *Note) Save() error {

	json, err := json.Marshal(&note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func Read() (Note, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return Note{}, err
	}

	var note Note
	err = json.Unmarshal(data, &note)
	return note, err
}