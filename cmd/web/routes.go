package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(secureHeaders(mux))
}
