package main

import (
	"Movies/pkg/forms"
	"Movies/pkg/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	movies, err := app.movies.Latest(10)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.render(w, r, "home.page.tmpl", &templateData{

		Movies: movies,
	})
}

func (app *application) horror(w http.ResponseWriter, r *http.Request) {
	movies, err := app.movies.GetMovieByGenre("horror")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "horror.page.tmpl", &templateData{Horror: movies})
}

func (app *application) comedy(w http.ResponseWriter, r *http.Request) {
	movies, err := app.movies.GetMovieByGenre("comedy")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "comedy.page.tmpl", &templateData{Comedy: movies})
}

func (app *application) drama(w http.ResponseWriter, r *http.Request) {
	movies, err := app.movies.GetMovieByGenre("drama")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "drama.page.tmpl", &templateData{Drama: movies})
}

func (app *application) sciFi(w http.ResponseWriter, r *http.Request) {
	movies, err := app.movies.GetMovieByGenre("sciFi")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "scifi.page.tmpl", &templateData{SciFi: movies})
}
func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "genre", "rating")
	form.MaxLength("title", 100)

	if !form.Valid() {
		app.render(w, r, "/", &templateData{Form: form})
		return
	}

	rating, err := strconv.ParseFloat(form.Get("rating"), 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	sessionTime, err := time.Parse("2006-01-02T15:04", form.Get("sessionTime"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	_, err = app.movies.Create(form.Get("title"), form.Get("genre"), rating, sessionTime)
	if errors.Is(err, models.ErrDuplicate) {
		app.clientError(w, http.StatusBadRequest)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Successfully created!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	id := r.PostForm.Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	genre := r.PostForm.Get("genre")
	ratingStr := r.PostForm.Get("rating")
	sessionTimeStr := r.PostForm.Get("sessionTime")

	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	sessionTime, err := time.Parse("2006-01-02T15:04", sessionTimeStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.movies.Update(objID, title, genre, rating, sessionTime)
	if errors.Is(err, models.ErrDuplicate) {
		app.clientError(w, http.StatusBadRequest)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("_id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.movies.Delete(objID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) contacts(w http.ResponseWriter, r *http.Request) {
	err := app.render(w, r, "contact.page.tmpl", nil)
	if err != nil {
		app.serverError(w, err)
	}
}
