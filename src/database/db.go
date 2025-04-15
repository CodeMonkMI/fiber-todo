package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=demo_app port=5433"
	c, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = c

	if err != nil {
		panic("Failed to connect to database")
	}

	return c

}

func GetDB() *gorm.DB {
	if db == nil {
		panic("Database not initialized")
	}
	return db
}
