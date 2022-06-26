package main

import (
	"fmt"
	"net/http"

	"github.com/captainmio/go-course/pkg/handlers"
)

const portNumber = ":8080"

// main is the entry point
func main() {

	// routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Start running the application in port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
