package middleware

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/jwt"
)

type Middleware struct {
	jwt jwt.JWTI
}

func NewMiddleware(jwt jwt.JWTI) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}
