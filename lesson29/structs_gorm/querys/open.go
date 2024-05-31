package querys

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	db, _ := gorm.Open(postgres.Open("postgres://postgres:pass@localhost:5432/new?sslmode=disable"))
	return db
}
