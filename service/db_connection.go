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

	return Connection(db_user, db_password, db_url, db_name)
}

func Connection(db_user string, db_password string, db_url string, db_name string) (db *gorm.DB, err error) {
	data := DBConfig{
		Username: db_user,
		Password: db_password,
		Url:      db_url,
		DBName:   db_name,
	}
	return data.ToClient(&gorm.Config{})
}

type DBConfig struct {
	Username string
	Password string
	Url      string
	DBName   string
}

func (config *DBConfig) ToDsn() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Url, config.DBName)
}

func (config *DBConfig) ToClient(gormCfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.ToDsn()), gormCfg)
}
