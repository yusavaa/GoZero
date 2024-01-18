package models

import (
	"GoZero/config"
	"GoZero/entities"
)

func GetAllArticle() []entities.Article {
	rows, err := config.DB.Query(`SELECT * FROM articles`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var articles []entities.Article

	for rows.Next() {
		var article entities.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Image, &article.Text); err != nil {
			panic(err)
		}

		articles = append(articles, article)
	}

	return articles
}

func GetTargetArticle(article_id int) entities.Article {
	rows := config.DB.QueryRow(`SELECT * FROM articles WHERE article_id = ?`, article_id)

	var article entities.Article
	if err := rows.Scan(&article.ID, &article.Title, &article.Image, &article.Text); err != nil {
		panic(err)
	}

	return article
}
