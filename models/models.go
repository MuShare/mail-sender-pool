package models

import (
	"fmt"
	"log"

	"gorm.io/gorm/schema"

	"github.com/MuShare/mail-sender-pool/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//Model base fields
type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfiguration.User,
		config.DatabaseConfiguration.Password,
		config.DatabaseConfiguration.Host,
		config.DatabaseConfiguration.Name)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed to connect to database. err: %v", err)
	}

	if err = migrateDatabase(); err != nil {
		log.Fatalf("failed to migrate database. err: %v", err)
	}
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	db, err := db.DB()
	if err != nil {
		defer db.Close()
	}
}
