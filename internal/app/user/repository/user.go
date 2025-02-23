package repository

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/entity"
	"gorm.io/gorm"
)

type UserMySQLItf interface {
	Create(user *entity.User) error
	Get(user *entity.User, dto dto.UserParam) error
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (r *UserMySQL) Create(user *entity.User) error {
	return r.db.Debug().Create(user).Error
}

func (r *UserMySQL) Get(user *entity.User, userParam dto.UserParam) error {
	return r.db.Debug().First(user, userParam).Error
}
