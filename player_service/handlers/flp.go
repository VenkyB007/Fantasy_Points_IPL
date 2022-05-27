package handlers

import (
	"net/http"
	"playerservice/db"
	"playerservice/models"

	"github.com/gin-gonic/gin"
)

func GetflpPlayers(c *gin.Context) {

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

	c.JSON(http.StatusOK, players)
	defer db.Close()

}

func GetflpScores(c *gin.Context) {

	db := db.OpenConnection()

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

	c.JSON(http.StatusOK, scores)

	defer db.Close()
}
