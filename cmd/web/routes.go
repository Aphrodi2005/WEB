package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := mux.NewRouter()

	mux.HandleFunc("/for-students", app.forStudents)
	mux.HandleFunc("/for-staff", app.forStaff)
	mux.HandleFunc("/for-applicants", app.forApplicants)
	mux.HandleFunc("/for-researches", app.forResearches)

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/createArticle", app.createArticle)
	mux.HandleFunc("/updateArticle", app.updateArticle).Methods("POST")
	mux.HandleFunc("/deleteArticle", app.deleteArticle).Methods("DELETE")
	mux.HandleFunc("/contacts", app.contacts)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	return mux
}
