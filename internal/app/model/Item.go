package model

type Item struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}
