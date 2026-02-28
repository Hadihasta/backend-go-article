package models

type Posts struct {
	ID       int `json:"id"`
	Title    int `json:"title"`
	Content  int `json:"content"`
	Category int `json:"category"`
	Status   int `json:"status"`
}