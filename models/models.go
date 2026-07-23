package models

import "time"

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	IsDone    bool   `json:"isdone"`
	AddedDate time.Time `json:"added_date"`
}


type RequestParams struct { 
		Title string `json:"title" binding:"required"`
		IsDone bool `json:"isdone" binding:"required"`
	}