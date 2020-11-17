package models

// Movie represents the model for an Movie
type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// Movies represents the model for an Movies
type Movies []Movie
