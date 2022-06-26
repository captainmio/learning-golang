package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/captainmio/go-course/pkg/config"
	"github.com/captainmio/go-course/pkg/handlers"
	"github.com/captainmio/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the entry point
func main() {

	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Start running the application in port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
