package bootstrap

import (
	"fmt"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/mysql"
)

func Start() error {

	config, err := LoadConfig()
	if err != nil {
		return err
	}

	err = mysql.Migrate(config.My)
	if err != nil {
		return err
	}

	MountRoutes(config)

	return config.Fb.Listen(fmt.Sprintf(":%d", config.En.AppPort))
}
