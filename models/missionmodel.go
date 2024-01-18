package models

import (
	"GoZero/config"
	"GoZero/entities"
	"net/http"
	"strings"
)

func GetAvailableMission(r *http.Request, id int) []entities.Mission {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	loginID, _ := session.Values["loginID"].(int)

	var missionIDs []string

	// Mengambil misi yang telah diselesaikan login user
	rowComplete, err := config.DB.Query(`SELECT mission_id FROM complete_missions WHERE user_id = ?`, loginID)
	if err == nil {
		for rowComplete.Next() {
			var missionID string
			if err := rowComplete.Scan(&missionID); err != nil {
				panic(err)
			}
			missionIDs = append(missionIDs, missionID)
		}
	}
	defer rowComplete.Close()

	// Mengambil misi yang belum diselesaikan login user
	query := "SELECT * FROM daily_missions"
	if len(missionIDs) > 0 {
		query += " WHERE mission_id NOT IN (" + strings.Join(missionIDs, ", ") + ")"
	}

	rowAvailable, err := config.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rowAvailable.Close()

	var missions []entities.Mission

	for rowAvailable.Next() {
		var mission entities.Mission
		if err := rowAvailable.Scan(&mission.ID, &mission.Description, &mission.Reward, &mission.Link); err != nil {
			panic(err)
		}

		missions = append(missions, mission)
	}

	return missions
}

func Complete(mission_id int, r *http.Request) bool {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	loginID, _ := session.Values["loginID"].(int)

	result, err := config.DB.Exec(`
		INSERT INTO complete_missions (user_id, mission_id)
		VALUE (?, ?)`,
		loginID, mission_id,
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
