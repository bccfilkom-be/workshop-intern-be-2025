package dto

import "github.com/google/uuid"

type RequestCreateProduct struct {
	Title       string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required"`
	Price       int64  `json:"price" validate:"required,min=1"`
	Stock       int8   `json:"stock" validate:"required,min=1"`
	PhotoUrl    string `json:"photo_url" validate:"required"`
}

type RequestUpdateProduct struct {
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title" validate:"omitempty,min=3,max=50"`
	Description string    `json:"description"`
	Price       int64     `json:"price" validate:"omitempty,min=1"`
	Stock       int8      `json:"stock" validate:"omitempty,min=1"`
	PhotoUrl    string    `json:"photo_url"`
}

type ResponseGetProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Stock       int8      `json:"stock"`
	PhotoUrl    string    `json:"photo_url"`
}
