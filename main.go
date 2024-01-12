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

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
