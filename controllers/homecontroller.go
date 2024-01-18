package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "home.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user": models.GetLoginUser(w, r),
		"photos": models.GetRandomPhotos(),
		"articles": models.GetAllArticle(),
	}

	tmpl.Execute(w, data)
}

