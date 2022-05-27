package handlers

import (
	"net/http"
	"playerservice/db"
	"playerservice/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePlayer godoc
// @Summary Create a new  Player
// @Description Create a new  Player with the input payload
// @Tags Players
// @Accept  json
// @Param Player body models.Player true "Player Details"
// @Success 200 {string} string "Successfully Posted!!"
// @Failure 400,500 {object} object
// @Router /player [post]
func CreatePlayer(c *gin.Context) {

	var p models.Player

	db := db.OpenConnection()

	error := c.Bind(&p)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	sqlstatement := `INSERT INTO playerinfo (name,id,team) VALUES($1, $2, $3)`

	_, err := db.Exec(sqlstatement, p.Name, p.Id, p.Team)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully Posted!!",
		})
	}

	defer db.Close()

}

// PostScores godoc
// @Summary Post scores
// @Description Post player scores with the input payload
// @Tags Scores
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param Score body models.Stats true "Player Scores"
// @Success 200 {string} string "Successfully Posted!!"
// @Router /player/{id}/score [post]
func PostScores(c *gin.Context) {

	db := db.OpenConnection()
	var s models.Stats
	err := c.Bind(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sqlstatement := `INSERT INTO scores (id, match, runs, wickets) VALUES($1, $2, $3, $4)`
	_, error := db.Exec(sqlstatement, id, s.Match, s.Runs, s.Wickets)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully Posted",
		})
	}

	defer db.Close()

}
