package main

import (
	"net/http"

	"github.com/emrahyurtlu/golang-bookings/cmd/pkg/config"
	"github.com/emrahyurtlu/golang-bookings/cmd/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	// Using middlewares.
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*",http.StripPrefix("/static", fileServer))
	
	return mux
}