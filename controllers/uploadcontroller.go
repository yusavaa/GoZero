package controllers

import (
	"GoZero/config"
	"GoZero/entities"
	"GoZero/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func ShowUploadPage(w http.ResponseWriter, r *http.Request) {
	if models.GetLoginUser(w, r).ID <= 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "page", "upload.html"))
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := os.CreateTemp("images", "image-*.jpg")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	tempFile.Write(fileBytes)
	image := filepath.Base(tempFile.Name())

	session, _ := config.Store.Get(r, config.SESSION_ID)

	var photo entities.Photo
	photo.Title = r.FormValue("title")
	photo.Image = image
	photo.Description = r.FormValue("description")
	photo.AuthorID = session.Values["loginID"].(int)

	if models.StorePhotoToDB(photo) {
		http.Redirect(w, r, "/gallery", http.StatusSeeOther)
	}

}
