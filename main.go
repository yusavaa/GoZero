package main

import (
	"GoZero/config"
	"GoZero/controllers"
	"net/http"
)

func main() {
	config.ConnectDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.ShowHomePage)

	mux.HandleFunc("/login", controllers.ShowLoginPage)
	mux.HandleFunc("/login/process", controllers.LoginForm)

	mux.HandleFunc("/register", controllers.ShowRegisterPage)
	mux.HandleFunc("/register/process", controllers.RegisterForm)

	mux.HandleFunc("/logout", config.DeleteSession)

	mux.HandleFunc("/gallery", controllers.ShowGalleryPage)
	mux.HandleFunc("/gallery/upload", controllers.ShowUploadPage)
	mux.HandleFunc("/gallery/upload/process", controllers.UploadFile)

	mux.HandleFunc("/quiz", controllers.ShowQuizPage)

	mux.HandleFunc("/article", controllers.ShowArticlePage)
	mux.HandleFunc("/article/read", controllers.ShowArticleDetilPage)

	mux.HandleFunc("/point", controllers.ShowPointPage)

	mux.HandleFunc("/banksampah", controllers.ShowBankMapPage)
	mux.HandleFunc("/bulkstore", controllers.ShowBulkMapPage)

	mux.HandleFunc("/mission", controllers.ShowMissionPage)

	mux.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
