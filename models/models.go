package models

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	IsDone    bool   `json:"isdone"`
	AddedDate string `json:"addeddate"`
}