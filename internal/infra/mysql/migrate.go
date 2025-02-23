package mysql

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Product{},
	); err != nil {
		return err
	}

	return nil
}
