package main

import (
	"errors"
	"net/http"
	"strconv"
	"tleukanov.net/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	articles, err := app.articles.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	categories, err := app.categories.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	contacts, err := app.contacts.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Pass the data to the home page template
	err = app.render(w, r, "home.page.tmpl", &templateData{
		Articles:   articles,
		Categories: categories,
		Contacts:   contacts,
	})
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	article, err := app.articles.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoArticle) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	err = app.render(w, r, "show.page.tmpl", &templateData{
		Article: article,
	})
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	category, err := app.categories.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoCategory) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	err = app.render(w, r, "category.page.tmpl", &templateData{
		Category: category,
		Articles: nil, // Add articles data if needed
		Contacts: nil, // Add contacts data if needed
	})
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	contact, err := app.contacts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoContact) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	err = app.render(w, r, "contact.page.tmpl", &templateData{
		Contact: contact,
	})
	if err != nil {
		app.serverError(w, err)
	}
}
func (app *application) createArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	err = app.articles.Create(title, content)
	if errors.Is(err, models.ErrDuplicate) {
		app.clientError(w, http.StatusBadRequest)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func (app *application) updateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := strconv.Atoi(r.PostForm.Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	err = app.articles.Update(id, title, content)
	if errors.Is(err, models.ErrDuplicate) {
		app.clientError(w, http.StatusBadRequest)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func (app *application) deleteArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PostForm.Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = app.articles.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}
