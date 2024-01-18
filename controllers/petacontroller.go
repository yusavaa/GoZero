package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowBulkMapPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "peta-bulkstore.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user": models.GetLoginUser(w, r),
	}

	tmpl.Execute(w, data)
}

func ShowBankMapPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "peta-banksampah.html"), path.Join("views", "partials", "navbar.html"), path.Join("views", "partials", "footer.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"user": models.GetLoginUser(w, r),
	}

	tmpl.Execute(w, data)
}
