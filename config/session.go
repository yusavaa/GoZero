package config

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const SESSION_ID = "go_auth"

var Store = sessions.NewCookieStore([]byte("fnaatfhisa"))

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, SESSION_ID)
	if err != nil {
		panic(err)
	}

	session.Options.MaxAge = -1
	session.Save(r, w)

	log.Println("Session deleted")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
