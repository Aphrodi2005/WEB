package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	mux.Get("/horror", dynamicMiddleware.ThenFunc(app.horror))
	mux.Get("/comedy", dynamicMiddleware.ThenFunc(app.comedy))
	mux.Get("/drama", dynamicMiddleware.ThenFunc(app.drama))
	mux.Get("/scifi", dynamicMiddleware.ThenFunc(app.sciFi))

	mux.Post("/createMovie", dynamicMiddleware.ThenFunc(app.createMovie))
	mux.Put("/updateMovie", dynamicMiddleware.ThenFunc(app.updateMovie))
	mux.Del("/deleteMovie", dynamicMiddleware.ThenFunc(app.deleteMovie))
	mux.Get("/contacts", dynamicMiddleware.ThenFunc(app.contacts))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}
