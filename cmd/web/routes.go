// routes.go
package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/article", app.showArticle)
	mux.HandleFunc("/contact", app.showContact)
	mux.HandleFunc("/category", app.showCategory)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}

// Add a new handler for articles
