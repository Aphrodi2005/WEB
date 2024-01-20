package main

import (
	"errors"
	"net/http"
	"strconv"
	"tleukanov.net/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	articles, err := app.articles.Latest(10)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Pass the data to the home page template
	err = app.render(w, r, "home.page.tmpl", &templateData{
		Articles: articles,
	})
	if err != nil {
		app.serverError(w, err)
	}
}
func (app *application) forStudents(w http.ResponseWriter, r *http.Request) {
	articles, err := app.articles.GetArticlesByCategory("For students")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "students.page.tmpl", &templateData{ForStudents: articles})
}

func (app *application) forStaff(w http.ResponseWriter, r *http.Request) {
	articles, err := app.articles.GetArticlesByCategory("For staff")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "staff.page.tmpl", &templateData{ForStaff: articles})
}

func (app *application) forApplicants(w http.ResponseWriter, r *http.Request) {
	articles, err := app.articles.GetArticlesByCategory("For applicants")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "applicants.page.tmpl", &templateData{ForApplicants: articles})
}

func (app *application) forResearches(w http.ResponseWriter, r *http.Request) {
	articles, err := app.articles.GetArticlesByCategory("For researches")
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "researches.page.tmpl", &templateData{ForResearches: articles})
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
	category := r.PostForm.Get("category")

	err = app.articles.Create(title, content, category)
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
	category := r.PostForm.Get("category")
	content := r.PostForm.Get("content")

	err = app.articles.Update(title, content, category, id)

	if errors.Is(err, models.ErrDuplicate) {
		app.clientError(w, http.StatusBadRequest)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) deleteArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.articles.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
