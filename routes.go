package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {

	// middlewares must be declared first

	// file server
	fileserver := http.FileServer(http.Dir("./public"))
	app.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileserver))

	// routes
	app.App.Routes.Get("/", app.Handlers.Home)

	return app.App.Routes
}
