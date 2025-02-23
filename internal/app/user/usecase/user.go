package usecase

import (
	"errors"

	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/user/repository"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/entity"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(register dto.Register) error
	Login(login dto.Login) (string, error)
}

type UserUsecase struct {
	userRepo repository.UserMySQLItf
	jwt      jwt.JWT
}

func NewUserUsecase(userRepo repository.UserMySQLItf, jwt jwt.JWT) UserUsecaseItf {
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

func (u *UserUsecase) Login(login dto.Login) (string, error) {
	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{Email: login.Email})
	if err != nil {
		return "", errors.New("invalid email or username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

    token, err := u.jwt.GenerateToken(user.ID, user.IsAdmin)
    if err != nil {
        return "", err
    }

    return token, nil
}
