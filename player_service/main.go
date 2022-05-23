package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Player struct {
	Name string `json:"name" binding:"required"`
	Id   int    `json:"id" binding:"required"`
	Team string `json:"team" binding:"required"`
}

type Score struct {
	Id      int    `json:"id" binding:"required"`
	Match   string `json:"match" binding:"required"`
	Runs    int    `json:"runs" binding:"required"`
	Wickets int    `json:"wickets" binding:"required"`
}

const (
	host     = "localhost"
	password = "venky"
	port     = 8008
	user     = "postgres"
	dbname   = "players"
)

func main() {
	r := gin.Default()

	// (i) : Post Player details
	r.POST("/player", func(c *gin.Context) {
		var p Player

		db := OpenConnection()

		c.Bind(&p)

		sqlstatement := `INSERT INTO players (name,id,team) VALUES($1, $2, $3)`

		db.QueryRow(sqlstatement, p.Name, p.Id, p.Team)
		// Here, query executing successfully, for getting panic err
		// thats why skipped the err

		c.String(200, "Successfully Posted!!")

		defer db.Close()
	})

	// (ii) : Post Player scores
	r.POST("/player/:id/score", func(c *gin.Context) {

		var s Score

		db := OpenConnection()

		c.Bind(&s)

		id := c.Params.ByName("id")

		sqlstatement := `INSERT INTO scores (id, match, runs, wickets) VALUES($1, $2, $3, $4)`

		db.QueryRow(sqlstatement, id, s.Match, s.Runs, s.Wickets)
		// Here, query executing successfully, for getting panic err
		// thats why skipped the err

		c.String(200, "Successfully Posted!!")

		defer db.Close()

	})

	// (iii) : Get all Player details from db
	r.GET("/players", func(c *gin.Context) {
		db := OpenConnection()

		rows, err := db.Query("SELECT * FROM players ")
		if err != nil {
			panic(err)
		}
		var players []Player
		for rows.Next() {
			var p Player
			rows.Scan(&p.Name, &p.Id, &p.Team)
			players = append(players, p)
		}

		c.JSON(200, gin.H{
			"players": players,
		})
		defer db.Close()
	})

	// (iv) Get all players scores
	r.GET("/players/scores", func(c *gin.Context) {
		db := OpenConnection()

		rows, err := db.Query("SELECT * FROM scores where id=1")
		if err != nil {
			panic(err)
		}
		var scores []Score
		for rows.Next() {
			var s Score
			rows.Scan(&s.Id, &s.Match, &s.Runs, &s.Wickets)
			scores = append(scores, s)
		}

		c.JSON(200, gin.H{
			"playerscores": scores,
		})

		defer db.Close()
	})

	r.Run("localhost:8989")
}

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("db connected successfully")
	return db
}
