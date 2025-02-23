package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTI interface{
	GenerateToken(userId uuid.UUID) (string, error)
}

type JWT struct{}

func NewJWT() JWTI {
	return &JWT{}
}

func (j *JWT) GenerateToken(userId uuid.UUID) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
