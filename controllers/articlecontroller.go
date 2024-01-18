package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowArticlePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "article.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user": models.GetLoginUser(w, r),
		"articles": models.GetAllArticle(),
	}

	tmpl.Execute(w, data)
}
