package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gohtml")
}

func renderTemplate(w http.ResponseWriter, gohtml string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parseTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// main is the entry point
func main() {

	// routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Start running the application in port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
