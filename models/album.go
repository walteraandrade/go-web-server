package models

import "time"

type Album struct {
    Title        string    `json:"title"`
    ReleaseDate time.Time `json:"releaseDate"`
    Label        string    `json:"label"`
    Artists      []string  `json:"artists"`
    Producer     string    `json:"producer"`
    Genres       []string  `json:"genres"`
    Image        string    `json:"image"`
    Rating       string    `json:"rating"`
    Review       string    `json:"review"`
    Tracks       []Track   `json:"tracks"`
}

type Track struct {
    Title  string `json:"title"`
    Number int    `json:"number"`
    Length int    `json:"length"` // In seconds
}
