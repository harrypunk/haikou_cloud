package service

import (
	"github.com/harrypunk/haikou_cloud/model"
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

func (client *Client) GetStudents(offset int, limit int) ([]model.Student, error) {
	var students []model.Student

	result := client.db.Offset(offset).Limit(limit).
		Select("id", "name", "gender", "age", "phone").
		Find(&students)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (client *Client) GetStudentDetail(id uint) (*model.Student, error) {
	var st model.Student
	result := client.db.
		Preload("School").
		First(&st, id)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &st, nil
}
