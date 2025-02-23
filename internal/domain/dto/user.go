package dto

import "github.com/google/uuid"

type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserParam struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
