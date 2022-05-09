package models

type Todos struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
