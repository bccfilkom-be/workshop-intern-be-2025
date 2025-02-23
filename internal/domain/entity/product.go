package entity

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int8      `gorm:"type:smallint;not null"`
	PhotoUrl    string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p *Product) ParseToDTO() dto.ResponseGetProduct {
	return dto.ResponseGetProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		PhotoUrl:    p.PhotoUrl,
	}
}
