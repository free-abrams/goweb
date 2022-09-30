package Models

import "time"

type Post struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
