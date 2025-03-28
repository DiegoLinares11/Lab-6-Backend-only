package models

type Match struct {
	ID          int    `json:"id"`
	HomeTeam    string `json:"homeTeam"`
	AwayTeam    string `json:"awayTeam"`
	MatchDate   string `json:"matchDate"`
	Goals       int    `json:"goals,omitempty"`
	YellowCards int    `json:"yellowCards,omitempty"`
	RedCards    int    `json:"redCards,omitempty"`
	ExtraTime   bool   `json:"extraTime,omitempty"`
}
