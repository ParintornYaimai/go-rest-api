package model

type BookModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
