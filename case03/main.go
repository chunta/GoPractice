package main

import (
	"bufio"
	"case03/note"
	"case03/structs"
	"fmt"
	"os"
	"strings"
	"time"
)

type nickName string

func (nickName *nickName) log() {
	fmt.Println(nickName)
}

func main() {
	// struct
	randomUser := structs.CreateARandomUser()
	randomUser.EchoUser()
	randomUser.EmptyUser()
	randomUser.EchoUser()

	// Json serialization
	newNote := note.Note{
		Title:     "New Title",
		Author:    "Sarah",
		CreatedAt: time.Now(),
		Content:   "0123",
	}
	fmt.Println(newNote)
	
	saveErr := newNote.Save()
	if saveErr != nil {
		fmt.Println("save failure")
	}

	readNote, jsonError := note.Read()
	if jsonError != nil {
		fmt.Println(jsonError)
	} else {
		fmt.Println("Read data:", readNote)
	}

	// read input
	fmt.Println("userInput:", getUserInput())

	// nickName custom type
	var nickName nickName = "im nick"
	nickName.log()

	// modify note using its pointer
	modifyNote(&newNote)
	fmt.Println(newNote.Content)
}

func getUserInput() string {
	fmt.Print("Input: ")
	// read long input
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Reading error")
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func modifyNote(note *note.Note) {
	note.Content = note.Content + "(Read)"
}