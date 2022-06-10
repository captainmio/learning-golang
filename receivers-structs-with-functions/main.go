package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Roger"

	myVar2 := myStruct{
		FirstName: "Joan",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
