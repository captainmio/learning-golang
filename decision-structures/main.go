package main

import "log"

func main() {
	var isTrue bool

	isTrue = false

	if isTrue == true {
		log.Println("isTrue is (true)", isTrue)
	} else {
		log.Println("isTrue is (false)", isTrue)
	}

	// =========================================

	cat := "Kitty"

	if cat == "Kitty" {
		log.Println("This Cat has a name", cat)
	} else {
		log.Println("Cat does not have name")
	}

	// =========================================

	myNum := 100
	booleanValue := false

	if myNum > 99 && booleanValue {
		log.Println("this line pass the condition")
	}

	// SWITCH STATEMENTS

	myAnimal := "fish1"

	switch myAnimal {
	case "cat":
		log.Println("this is a cat")
	case "dog":
		log.Println("this is a dog")
	case "fish":
		log.Println("this is a fish")
	default:
		log.Println("this isn't an animal")
	}
}
