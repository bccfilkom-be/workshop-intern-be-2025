package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTI interface{
	GenerateToken(userId uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
}

type JWT struct{
	SecretKey string
	ExpiredTime time.Time
}


func NewJWT() *JWT {
	secretKey := "SECRET"
	expiredTime := time.Now().Add(time.Hour * 2)

	return &JWT{
		SecretKey: secretKey,
		ExpiredTime: expiredTime,
	}
}

type Claims struct {
	UserId uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(userId uuid.UUID, isAdmin bool) (string, error) {
	claim := Claims{
		UserId: userId,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, error) {
	var claim Claims
	var id uuid.UUID

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return id, err
	}

	if !token.Valid {
		return id, err
	}

	id = claim.UserId

	return id, nil
}
