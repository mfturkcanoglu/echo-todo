package entity

import "time"

type Todo struct {
	Id        int64     `json:"id" query:"id"`
	Text      string    `json:"text" query:"text"`
	CreatedAt time.Time `json:"created_at" query:"created_at"`
	Deleted   bool      `json:"deleted" query:"deleted"`
}
