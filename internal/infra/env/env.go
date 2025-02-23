package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Env struct {
	AppName string `env:"APP_NAME,required"`
	AppPort int    `env:"APP_PORT,required"`

	DatabaseHost     string `env:"DB_HOST,required"`
	DatabasePort     int    `env:"DB_PORT,required"`
	DatabaseUsername string `env:"DB_USERNAME,required"`
	DatabasePassword string `env:"DB_PASSWORD,required"`
	DatabaseName     string `env:"DB_NAME,required"`

	JwtSecret string `env:"JWT_SECRET,required"`
	JwtExpired int `env:"JWT_EXPIRED,required"`
}

func New() (*Env, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := new(Env)

	err = env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
