package controllers

import (
	"GoZero/models"
	"html/template"
	"net/http"
	"path"
)

func ShowGalleryPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "page", "gallery.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"photos": models.GetAllPhoto(),
	}

	tmpl.Execute(w, data)
}


