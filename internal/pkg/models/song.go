package models

type Song struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
