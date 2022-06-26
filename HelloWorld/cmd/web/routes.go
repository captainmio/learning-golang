package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/captainmio/go-course/pkg/config"
	"github.com/captainmio/go-course/pkg/handlers"
)

// Routes applying Pat router
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
