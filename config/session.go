package config

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(""))

func GetSession(r *http.Request, sessionName string) *sessions.Session {
	session, err := Store.Get(r, sessionName)
	if err != nil {
		panic(err)
	}

	return session
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "loginID")
	if err != nil {
		panic(err)
	}

	session.Options.MaxAge = -1
	session.Save(r, w)

	log.Println("Session deleted")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
