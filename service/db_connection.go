package service

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func EnvConnection() (db *gorm.DB, err error) {
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_url := os.Getenv("DB_URL")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_url, db_name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
