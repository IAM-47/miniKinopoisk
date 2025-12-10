package models

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Producer    string `json:"producer"`
	Director    string `json:"director"`
	ReleaseYear int    `json:"release_year"`
}
