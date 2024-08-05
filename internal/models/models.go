package models

type Books struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Published_at int    `json:"published_at"`
}
