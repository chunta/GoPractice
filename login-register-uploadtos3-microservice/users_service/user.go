package main

import (
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
