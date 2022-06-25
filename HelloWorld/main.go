package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	mySum := addValues(2, 5)
	fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2 + 5 is %d", mySum))
}

// addValues is a method to calculate 2 integer and returns the sum
func addValues(x int, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float64 = 100.0
	var y float64 = 0.0
	f, err := divideValues(x, y)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
	} else {
		fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", x, y, f))
	}

}

func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	var result float64
	result = x / y
	return result, nil
}

// main is the entry point
func main() {

	// routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Start running the application in port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
