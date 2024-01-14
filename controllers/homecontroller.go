package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "home.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user": models.GetLoginUser(w, r),
	}

	tmpl.Execute(w, data)
}
