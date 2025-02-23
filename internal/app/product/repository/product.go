package repository

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	Create(product *entity.Product) error
	GetAll(products *[]entity.Product) error
	GetSpecific(product *entity.Product) error
	Delete(product *entity.Product) error
	Updates(product *entity.Product) error
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (r *ProductMySQL) Create(product *entity.Product) error {
	return r.db.Debug().Create(product).Error
}

func (r *ProductMySQL) GetAll(products *[]entity.Product) error {
	return r.db.Debug().Find(products).Error
}

func (r *ProductMySQL) GetSpecific(product *entity.Product) error {
	return r.db.Debug().First(product).Error
}

func (r *ProductMySQL) Delete(product *entity.Product) error {
	return r.db.Debug().Delete(product).Error
}

func (r *ProductMySQL) Updates(product *entity.Product) error {
	return r.db.Debug().Updates(product).Error
}
