package usecase

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/user/repository"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(register dto.Register) error
}

type UserUsecase struct {
	userRepo repository.UserMySQLItf
}

func NewUserUsecase(userRepo repository.UserMySQLItf) UserUsecaseItf {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Register(register dto.Register) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

    user := entity.User{
        Name:     register.Name,
        Email:    register.Email,
        Password: string(hashedPassword),
    }

    err = u.userRepo.Create(&user)
    
    return err
}
