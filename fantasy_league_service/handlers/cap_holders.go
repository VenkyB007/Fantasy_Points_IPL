package handlers

import (
	"encoding/json"
	"fantasy_league_service/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCapHolders godoc
// @Summary Get CapHolders
// @Description Get Cap Holders
// @Tags Fantasy_League_Service
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router /cap-holders [get]
func CapHodler(c *gin.Context) {

	resp, err := http.Get("http://localhost:8989/players/flpscores")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("No response from request")
	}
	var scores []models.Score
	json.Unmarshal(body, &scores)

	var ids []int
	for i := range scores {
		ids = append(ids, scores[i].Id)
	}

	rd := RemoveDuplicate(ids)
	var runs []int
	var wickets []int
	for _, v := range rd {
		rsum := 0
		wsum := 0
		for j := range scores {
			if v == scores[j].Id {
				rsum = rsum + scores[j].Runs
				wsum = wsum + scores[j].Wickets
			}
		}
		runs = append(runs, rsum)
		wickets = append(wickets, wsum)
	}
	// fmt.Println(rd)
	// fmt.Println(runs)
	// fmt.Println(wickets)
	var dpruns []int
	var dpwickets []int
	for i := 0; i < len(runs); i++ {
		dpruns = append(dpruns, runs[i])
		dpwickets = append(dpwickets, wickets[i])
	}
	//fmt.Println("dp runs", dpruns)
	//sorting
	for i := 0; i < len(dpruns)-1; i++ {
		if dpruns[i] > dpruns[i+1] {
			dpruns[i], dpruns[i+1] = dpruns[i+1], dpruns[i]
		}
		if dpwickets[i] > dpwickets[i+1] {
			dpwickets[i], dpwickets[i+1] = dpwickets[i+1], dpwickets[i]
		}
	}
	// fmt.Println("sort", dpruns)
	// fmt.Println("runs", runs)
	rmax := dpruns[len(dpruns)-1]
	wmax := dpwickets[len(dpwickets)-1]

	//fmt.Println(rmax)
	var O_cap_id_index int
	var P_cap_id_index int

	for i := 0; i < len(runs); i++ {
		if runs[i] == rmax {
			O_cap_id_index = rd[i]
		}
		if wickets[i] == wmax {
			P_cap_id_index = rd[i]
		}
	}
	//fmt.Println("O_capholder id index", O_cap_id_index)

	orange_cap_holder := OrangeCap(O_cap_id_index)
	Purple_cap_holder := PurpleCap(P_cap_id_index)

	c.JSON(http.StatusOK, gin.H{
		"OrangeCap": orange_cap_holder,
		"PurpleCap": Purple_cap_holder,
	})

}

func OrangeCap(id int) (name string) {
	//Getting player Info
	resp, err := http.Get("http://localhost:8989/players/flp")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("No response from request")
	}

	var players []models.Player
	json.Unmarshal(body, &players)

	//fmt.Println(players)
	for i := 0; i < len(players); i++ {
		if players[i].Id == id {

			return players[i].Name

		}
	}
	return
}

func PurpleCap(id int) (name string) {
	//Getting player Info
	resp, err := http.Get("http://localhost:8989/players/flp")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("No response from request")
	}

	var players []models.Player
	json.Unmarshal(body, &players)

	//fmt.Println(players)
	for i := 0; i < len(players); i++ {
		if players[i].Id == id {

			return players[i].Name

		}
	}
	return
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
