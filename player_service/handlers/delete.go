package handlers

import (
	"net/http"
	"playerservice/db"

	"github.com/gin-gonic/gin"
)

// DeletePlayer godoc
// @Summary Delete details of Player based on player id
// @Description Delete details of Player
// @Tags Players
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} string "Successfully Deleted!!"
// @Router /player/{id} [delete]
func DeletePlayer(c *gin.Context) {
	db := db.OpenConnection()

	id := c.Param("id")
	sqlstatement := `delete from playerinfo where id=$1`

	_, err := db.Exec(sqlstatement, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted successfully",
		})
	}
	defer db.Close()
}

// DeletePlayerScores godoc
// @Summary Delete Scores of Player based on player id
// @Description Delete Scores of Player
// @Tags Scores
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} string "Successfully Deleted!!"
// @Router /player/{id}/score [delete]
func DeletePlayerScores(c *gin.Context) {
	db := db.OpenConnection()

	id := c.Param("id")
	sqlstatement := `delete from scores where id=$1`

	_, err := db.Exec(sqlstatement, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted successfully",
		})
	}
	defer db.Close()
}
