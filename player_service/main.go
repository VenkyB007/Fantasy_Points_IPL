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
	Match   string `json:"match" binding:"required"`
	Runs    int    `json:"runs" binding:"required"`
	Wickets int    `json:"wickets" binding:"required"`
}

type Id struct {
	Id    int
	Score []Score
}

const (
	host     = "localhost"
	password = "1234"
	port     = 5432
	user     = "postgres"
	dbname   = "player"
)

func main() {
	r := gin.Default()

	r.POST("/player", func(c *gin.Context) {
		var p Player

		db := OpenConnection()

		c.Bind(&p)

		//fmt.Printf("api params %s %d %s", p.Name, p.Id, p.Team)
		sqlstatement := `INSERT INTO playerinfo (name,id,team) VALUES($1, $2, $3)`

		db.QueryRow(sqlstatement, p.Name, p.Id, p.Team)
		// Here, query executing successfully, for getting panic err
		// thats why skipped the err

		c.String(200, "Successfully Posted!!")

		defer db.Close()
	})

	r.GET("/players", func(c *gin.Context) {
		db := OpenConnection()

		rows, err := db.Query("SELECT * FROM playerinfo")
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

	r.POST("/player/:id/score", func(c *gin.Context) {

		var s Score

		db := OpenConnection()

		c.Bind(&s)

		sqlstatement := `INSERT INTO scores (match, runs, wickets) VALUES($1, $2, $3)`

		var id int
		db.QueryRow(sqlstatement, s.Match, s.Runs, s.Wickets).Scan(&id)
		// Here, query executing successfully, for getting panic err
		// thats why skipped the err

		c.String(200, "Successfully Posted!!")

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