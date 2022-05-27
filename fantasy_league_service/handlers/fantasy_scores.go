package handlers

import (
	"encoding/json"
	"fantasy_league_service/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFantasy Scores godoc
// @Summary Get Fantasy Scores
// @Description Get Fantasy Scores
// @Tags Fantasy_League_Service
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router /fantasy-scores [get]
func FantasyScores(c *gin.Context) {
	//Get Players
	response, err := http.Get("http://localhost:8989/players/flp")
	if err != nil {
		fmt.Println("No response from Request")
	}
	defer response.Body.Close()
	playerbody, err := ioutil.ReadAll(response.Body) // response body is []byte
	if err != nil {
		fmt.Println("No response from request")
	}

	var players []models.Player
	json.Unmarshal(playerbody, &players)

	//Get Scores
	resp, err := http.Get("http://localhost:8989/players/flpscores")
	if err != nil {
		fmt.Println("No response from Request")
	}
	defer resp.Body.Close()
	scoresbody, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("No response from request")
	}

	var scores []models.Score
	json.Unmarshal(scoresbody, &scores)
	//fmt.Println(scores)

	var rd []int
	for i := range scores {
		rd = append(rd, scores[i].Id)
	}

	ids := flpRemoveDuplicate(rd)
	//fmt.Println(ids)

	scoremap := map[int]int{}

	for _, v := range ids {
		scoresum := 0
		for i := range scores {
			if scores[i].Id == v {
				//runs score
				if scores[i].Runs >= 100 {
					scoresum = scoresum + 170
				} else if scores[i].Runs >= 50 {
					scoresum = scoresum + 70
				} else if scores[i].Runs >= 30 {
					scoresum = scoresum + 20
				}

				//wickets score
				if scores[i].Wickets >= 5 {
					scoresum = scoresum + 50
				}
				scoresum = scoresum + scores[i].Wickets*10
			}
		}
		scoremap[v] = scoresum
	}

	fmt.Println("scores map :", scoremap)

	var fantasyScores []models.FantasyScores
	var fs models.FantasyScores
	for i := range players {
		for k, v := range scoremap {
			if players[i].Id == k {
				fs.Name = players[i].Name
				fs.FantasyScore = v
			}
		}
		fantasyScores = append(fantasyScores, fs)
	}

	for i := range fantasyScores {
		for j := i + 1; j < len(fantasyScores); j++ {
			if fantasyScores[i].FantasyScore < fantasyScores[j].FantasyScore {
				fantasyScores[i], fantasyScores[j] = fantasyScores[j], fantasyScores[i]
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"fantasyScores": fantasyScores,
	})

}

func flpRemoveDuplicate(a []int) (r []int) {
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
