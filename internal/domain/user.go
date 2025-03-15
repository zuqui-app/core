package domain

import "time"

type User struct {
	Id        string    `json:"id,omitempty"`
	Email     string    `json:"email"        validate:"required"`
	Username  string    `json:"username"     validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
