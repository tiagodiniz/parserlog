package model

type BodyCount struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
	FilePath   string         `json:"-"`
	Error      error          `json:"-"`
}

