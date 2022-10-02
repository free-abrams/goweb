package Rule

import "time"

type Rule struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Rule      string    `json:"rule"`
	IsMenu    int       `json:"is_menu"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
