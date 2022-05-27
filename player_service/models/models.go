package models

type Player struct {
	Name string `json:"name" `
	Id   int    `json:"id" `
	Team string `json:"team" `
}

type Score struct {
	Id      int `json:"id"`
	Match   int `json:"match" `
	Runs    int `json:"runs" `
	Wickets int `json:"wickets" `
}

type Stats struct {
	Match   int `json:"match" `
	Runs    int `json:"runs" `
	Wickets int `json:"wickets" `
}

type PlayerScores struct {
	Id    int
	Stats []Stats
}
