package bootstrap

import (
	"fmt"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/env"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/fiber"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/mysql"
	"github.com/go-playground/validator/v10"
	fb "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Config struct {
	Fb *fb.App
	En *env.Env
	My *gorm.DB
	Va *validator.Validate
}

func LoadConfig() (*Config, error) {
	_env, err := env.New()
	if err != nil {
		return nil, err
	}

	_mysql, err := mysql.New(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			_env.DatabaseUsername,
			_env.DatabasePassword,
			_env.DatabaseHost,
			_env.DatabasePort,
			_env.DatabaseName),
	)
	if err != nil {
		return nil, err
	}

	_fiber := fiber.New()
	va := validator.New()

	return &Config{
		Fb: _fiber,
		En: _env,
		My: _mysql,
		Va: va,
	}, nil
}
