package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	Data DBConfig
	db   *gorm.DB
}

func NewClient(debug bool, data DBConfig) (*Client, error) {
	cfg := gorm.Config{}
	if debug {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := data.ToClient(&cfg)
	if err != nil {
		return nil, err
	}
	return &Client{
		Data: data,
		db:   db,
	}, nil
}

func (client *Client) GetStudents(start int, end int) {

}
