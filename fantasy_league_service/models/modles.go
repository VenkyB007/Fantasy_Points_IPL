package models

type Score struct {
	Id      int `json:"id"`
	Match   int `json:"match"`
	Runs    int `json:"runs"`
	Wickets int `json:"wickets"`
}

type Player struct {
	Name string `json:"name" `
	Id   int    `json:"id" `
	Team string `json:"team" `
}

type FantasyScores struct {
	Name         string `json:"name"`
	FantasyScore int    `json:"fantasyScore"`
}
