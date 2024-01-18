package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func ShowArticleDetilPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "article-detil.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user":    models.GetLoginUser(w, r),
		"article": models.GetTargetArticle(id),
	}

	if models.GetLoginUser(w, r).ID <= 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else {
		models.Complete(1, r)
		models.UpdatePoint(100, w, r)
	}

	tmpl.Execute(w, data)
}
