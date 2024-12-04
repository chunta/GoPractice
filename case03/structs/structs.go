package structs

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type User struct {
	FirstName string
	LastName  string
	UpdateAt  time.Time
}

type Admin struct {
	Email string
	User
}

func CreateARandomUser() User {
	user := User{
		FirstName: randomdata.FirstName(rand.Intn(1)),
		LastName:  randomdata.LastName(),
		UpdateAt:  time.Now(),
	}
	return user
}

func (user *User) EmptyUser() {
	user.FirstName = ""
	user.LastName = ""
	user.UpdateAt = time.Now()
}

func (user *User) EchoUser() {
	fmt.Printf("firstName: %s,\nlastName: %s,\nupdate: %v)\n\n",
		user.FirstName,
		user.LastName,
		user.UpdateAt)
}
