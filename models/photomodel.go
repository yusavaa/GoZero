package models

import (
	"GoZero/config"
	"GoZero/entities"
)

func GetAllPhoto() []entities.Photo {
	rows, err := config.DB.Query(`SELECT * FROM photos`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var photos []entities.Photo

	for rows.Next() {
		var photo entities.Photo
		if err := rows.Scan(&photo.ID, &photo.Title, &photo.Image, &photo.Description, &photo.AuthorID); err != nil {
			panic(err)
		}

		photos = append(photos, photo)
	}

	return photos
}

func StorePhotoToDB(photo entities.Photo) bool {
	result, err := config.DB.Exec(`
		INSERT INTO photos (title, image, description, author_id)
		VALUE (?, ?, ?, ?)`,
		photo.Title, photo.Image, photo.Description, photo.AuthorID,
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