package models

import (
	"GoZero/config"
	"GoZero/entities"
	"net/http"
)

func GetLoginUser(w http.ResponseWriter, r *http.Request) entities.User {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	loginID, _ := session.Values["loginID"].(int)

	rows := config.DB.QueryRow(`SELECT * FROM users WHERE user_id = ?`, loginID)

	var user entities.User
	rows.Scan(&user.ID, &user.Username, &user.Password, &user.Point)

	return user
}

func Login(username string, password string, w http.ResponseWriter, r *http.Request) bool {
	var flag bool

	rows := config.DB.QueryRow(`SELECT * FROM users WHERE username = ?`, username)

	var user entities.User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Point); err != nil {
		flag = false
	}

	if user.Password == password {
		session, _ := config.Store.Get(r, config.SESSION_ID)

		session.Values["loginID"] = user.ID
		session.Save(r, w)

		flag = true
	}

	return flag
}

func Create(user entities.User) bool {
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

	return lastInsertId > 0
}
