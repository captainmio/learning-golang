package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		// log.Println(i)
	}

	animals := []string{"dog", "goat", "cat", "bird"}

	// with key
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// remove key
	for _, animal := range animals {
		log.Println(animal)
	}

	// ====================================

	things := make(map[string]string)

	things["one"] = "Pocket"
	things["two"] = "Phone"

	for thingType, thing := range things {
		log.Println(thingType, thing)
	}

	// ====================================

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{
		"John", "Smith", "john@smith.com", 30,
	})

	users = append(users, User{
		"Mary", "Brown", "mary@smith.com", 30,
	})

	users = append(users, User{
		"Cory", "Batum", "cory@smith.com", 22,
	})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}
