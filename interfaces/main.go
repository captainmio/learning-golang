package main

import "fmt"

// define interface
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Roger",
		Breed: "Poodle",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Chuck",
		Color:         "Black",
		NumberOfTeeth: 30,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Boohh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
