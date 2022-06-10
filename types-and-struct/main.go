package main

// topic for global variable and variable showing
// import "log"

// var s = "seven"

// func main() {
// 	var s2 = "six"
// 	log.Println("s is", s)
// 	log.Println("s2 is", s2)

// 	saySomething("xxx")
// }

// func saySomething(s3 string) (string, string) {
// 	log.Println("s from the saySomething func is", s)
// 	return s3, "World"
// }

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "22331132323",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate: ", user.BirthDate)

}
