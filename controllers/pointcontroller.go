package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowPointPage(w http.ResponseWriter, r *http.Request) {
	if models.GetLoginUser(w, r).ID <= 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "page", "point.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user":    models.GetLoginUser(w, r),
	}

	tmpl.Execute(w, data)
}
