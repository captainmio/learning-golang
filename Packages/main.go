package main

import (
	"log"

	"github.com/captainmio/samplePackage/helpers"
)

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(100)
	intChan <- randomNumber
}

func main() {
	log.Println("Sample Package")
	var myVar helpers.SomeType

	myVar.TypeName = "Something Naming"

	log.Println(myVar.TypeName)

	intChan := make(chan int)
	defer close(intChan)

	// this execute the function as a go routine
	go CalculateValue(intChan)

	num := <-intChan
	log.Print(num)
}

// use the command "go mod init"; example command is go mod init github.com/captainmio/samplePackage
