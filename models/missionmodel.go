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
	var result string
	var query string

	// mengambil misi yang telah diselesaikan login user
	rowComplete, err := config.DB.Query(`SELECT mission_id FROM complete_missions WHERE user_id = ?`, loginID)
	if err == nil {
		for rowComplete.Next() {
			var missionID string
			if err := rowComplete.Scan(&missionID); err != nil {
				panic(err)
			}
			missionIDs = append(missionIDs, missionID)
			result = strings.Join(missionIDs, ", ")
			// mengambil misi yang belum diselesaikan login user
			query = "SELECT * FROM daily_missions WHERE mission_id NOT IN (" + result + ")"
		}
	}
	defer rowComplete.Close()

	if query == "" {
		query = "SELECT * FROM daily_missions"
	}

	rowAvailable, err := config.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rowAvailable.Close()

	var missions []entities.Mission

	for rowAvailable.Next() {
		var mission entities.Mission
		if err := rowAvailable.Scan(&mission.ID, &mission.Description, &mission.Reward); err != nil {
			panic(err)
		}

		missions = append(missions, mission)
	}

	return missions
}
