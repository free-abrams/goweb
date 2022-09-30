package Models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Post      []Post    `json:"post"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SelectOne(user User) (*User, error) {
	result := DbConn.First(&user)
	if result.RowsAffected > 0 {
		post := Post{}
		result := DbConn.First(&post, "user_id = ?", user.Id)
		if result.RowsAffected > 0 {
			user.Post = append(user.Post, post)
		}
		return &user, nil
	}

	var err error
	return nil, err
}
