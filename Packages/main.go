package main

import (
	"log"

	"github.com/captainmio/samplePackage/helpers"
)

func main() {
	log.Println("Sample Package")
	var myVar helpers.SomeType

	myVar.TypeName = "Something Naming"

	log.Println(myVar.TypeName)
}

// use the command "go mod init"; example command is go mod init github.com/captainmio/samplePackage
