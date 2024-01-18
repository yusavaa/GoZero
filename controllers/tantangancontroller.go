package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowMissionPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "tantangan.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user":   models.GetLoginUser(w, r),
		"missions": models.GetAvailableMission(r, 1),
	}

	tmpl.Execute(w, data)
}
