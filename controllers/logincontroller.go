package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "login.html"))
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}

func LoginForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if models.Login(username, password, w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
