package controllers

import (
	"GoZero/config"
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

	isLogin := false

	if session := config.GetSession(r, "loginID"); session.Values["loginID"].(int) > 0 {
		isLogin = true
	}

	userData := models.GetLoginUser(w, r)

	data := map[string]interface{}{
		"IsLogin": isLogin,
		"User":    userData,
	}

	tmpl.Execute(w, data)
}
