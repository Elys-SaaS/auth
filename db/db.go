package db

import (
	"fmt"

	"github.com/Elys-SaaS/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = "host=localhost user=postgres password=mysecretpassword dbname=test-pg port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func New() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}

func DropTestDB() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	if err := db.Exec("DROP DATABASE gorm_test").Error; err != nil {
		return err
	}

	return nil
}

// TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
