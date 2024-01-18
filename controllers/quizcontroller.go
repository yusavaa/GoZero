package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func ShowQuizPage(w http.ResponseWriter, r *http.Request) {
	if models.GetLoginUser(w, r).ID <= 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "page", "quiz.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user": models.GetLoginUser(w, r),
		"quiz": models.GetAllQuiz(id),
	}

	models.Complete(2, r)
	models.UpdatePoint(200, w, r)

	tmpl.Execute(w, data)
}
