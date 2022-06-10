package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// ---- MAPS
	myMap := make(map[string]string) //short hand way to create a map

	myMap["dog"] = "Luigi"
	myMap["other-dog"] = "Casey"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// replacing value
	myMap["dog"] = "Buck"
	// log.Println(myMap["dog"])

	// For Int
	myMapInt := make(map[string]int)

	myMapInt["first"] = 1
	myMapInt["second"] = 2

	// log.Println(myMapInt["first"])
	// log.Println(myMapInt["second"])

	// create your own data type using struct
	myMapOwnType := make(map[string]User)

	me := User{
		FirstName: "Ricky",
		LastName:  "Valdueza",
	}

	myMapOwnType["me"] = me

	// log.Println(myMapOwnType["me"].FirstName)

	// ---- SLICES
	var lists []string
	lists = append(lists, "Boris")
	lists = append(lists, "Pablo")
	lists = append(lists, "Anne")

	// log.Println(lists)

	var listsInt []int
	listsInt = append(listsInt, 4)
	listsInt = append(listsInt, 1)
	listsInt = append(listsInt, 3)

	// log.Println(listsInt)

	sort.Ints(listsInt) // do sorting of Integer value

	// log.Println(listsInt)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	names := []string{"one", "two", "three"}

	log.Println(numbers)

	// display only with range
	log.Println(numbers[6:9])

	log.Println(names)
	log.Println(names[0:2])
}
