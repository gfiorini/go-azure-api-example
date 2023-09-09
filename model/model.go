package model

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Info struct {
	Version string `json:"version"`
}

type Leaderboard struct {
	Scores []Score `json:"scores"`
}
type Score struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
