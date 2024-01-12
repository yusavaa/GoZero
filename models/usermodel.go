package models

import (
	"GoZero/config"
	"GoZero/entities"
	"database/sql"
	"net/http"
)

func GetLoginUser(w http.ResponseWriter, r *http.Request) *entities.User {
	session := config.GetSession(r, "loginID")

	loginID, ok := session.Values["loginID"].(int)
	if !ok {
		loginID = 0
	}

	rows := config.DB.QueryRow(`SELECT * FROM users WHERE user_id = ?`, loginID)

	var user entities.User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user
}

func Login(username string, password string, w http.ResponseWriter, r *http.Request) bool {
	login := false

	rows := config.DB.QueryRow(`SELECT * FROM users WHERE username = ?`, username)

	var user entities.User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		panic(err)
	}

	if user.Password == password {
		session := config.GetSession(r, "loginID")

		session.Values["loginID"] = user.ID
		session.Save(r, w)

		login = true
	}

	return login
}

func Create(user entities.User) int {
	result, err := config.DB.Exec(`
		INSERT INTO users (username, password)
		VALUE (?, ?)`,
		user.Username, user.Password,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(lastInsertId)
}
