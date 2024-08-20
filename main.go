package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Miracle",
		LastName:    "Chris",
		PhoneNumber: "0700000000",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber)
}
