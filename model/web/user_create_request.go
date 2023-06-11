package web

import "time"

type CreateUserRequest struct {
	Id        string    `json:"id"`
	Firstname string    `json:"first_name" validate:"required,min=3"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
