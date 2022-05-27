package handlers

import (
	"net/http"
	"playerservice/db"
	"playerservice/models"

	"github.com/gin-gonic/gin"
)

// GetPlayers godoc
// @Summary Get details of all Players from the database
// @Description Get details of all Players
// @Tags Players
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Player
// @Router /players [get]
func GetPlayers(c *gin.Context) {

	db := db.OpenConnection()

	rows, err := db.Query("SELECT * FROM playerinfo ")
	if err != nil {
		panic(err)
	}
	var players []models.Player
	for rows.Next() {
		var p models.Player
		rows.Scan(&p.Id, &p.Name, &p.Team)
		players = append(players, p)
	}

	c.JSON(http.StatusOK, gin.H{
		"players": players,
	})
	defer db.Close()

}

// GetPlayerScores godoc
// @Summary Get Scores of all Players from the database
// @Description Get Scores of all Players
// @Tags Scores
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Score
// @Router /players/scores [get]
func GetScores(c *gin.Context) {

	db := db.OpenConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM scores")
	if err != nil {
		panic(err)
	}
	var scores []models.Score
	for rows.Next() {
		var s models.Score
		rows.Scan(&s.Id, &s.Match, &s.Runs, &s.Wickets)
		scores = append(scores, s)
	}

	var rd []int
	for i := range scores {
		rd = append(rd, scores[i].Id)
	}

	ids := RemoveDuplicate(rd)

	for i := range ids {
		for j := i + 1; j < len(ids); j++ {
			if ids[i] > ids[j] {
				ids[i], ids[j] = ids[j], ids[i]
			}
		}
	}

	var playerscores []models.PlayerScores
	for _, v := range ids {
		var stats []models.Stats
		for k := range scores {
			if v == scores[k].Id {
				var stat models.Stats
				stat.Match = scores[k].Match
				stat.Runs = scores[k].Runs
				stat.Wickets = scores[k].Wickets

				stats = append(stats, stat)
			}
		}
		var ps models.PlayerScores
		ps.Id = v
		ps.Stats = stats
		playerscores = append(playerscores, ps)
	}

	c.JSON(http.StatusOK, gin.H{
		"PlayerScores": playerscores,
	})
}

func RemoveDuplicate(a []int) (r []int) {
	m := map[int]int{}
	for _, v := range a {
		m[v]++
	}
	for _, v := range a {
		if m[v] == 1 {
			r = append(r, v)
		}
		m[v]--
	}
	return
}
