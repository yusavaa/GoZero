package controllers

import (
	"GoZero/config"
	"GoZero/entities"
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "register.html"))
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}

func RegisterForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	var user entities.User
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	loginID := models.Create(user)

	session, err := config.Store.Get(r, "loginID")
	if err != nil {
		panic(err)
	}

	session.Values["loginID"] = loginID
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
