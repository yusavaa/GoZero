package models

import (
	"GoZero/config"
	"GoZero/entities"
)

func GetAllQuiz(article_id int) []entities.Quiz {
	rows, err := config.DB.Query(`SELECT quiz_id, question, opt1, opt2, opt3, answer FROM quizzes WHERE article_id = ?`, article_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var quizzes []entities.Quiz

	for rows.Next() {
		var quiz entities.Quiz
		if err := rows.Scan(&quiz.ID, &quiz.Question, &quiz.Option1, &quiz.Option2, &quiz.Option3, &quiz.Answer); err != nil {
			panic(err)
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes
}
