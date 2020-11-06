package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func migrateDatabase() error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "V1",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&SMTPAccount{})
			},
		},
	})
	return m.Migrate()
}
