package main

import (
	"net/http"

	"github.com/captainmio/go-course/pkg/config"
	"github.com/captainmio/go-course/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Routes applying Pat router
// func routes(app *config.AppConfig) http.Handler {
// 	mux := pat.New()

// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

// 	return mux
// }

// Routes applying Chi router
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
