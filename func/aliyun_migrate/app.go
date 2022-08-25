package main

import (
	"context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	m "github.com/harrypunk/haikou_cloud/model"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func HandleRequest(ctx context.Context, event StructEvent) (*string, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&m.Student{},
		&m.Parent{},
		&m.Family{},
		&m.School{},
		&m.Grade{},
		&m.Course{},
		&m.Teacher{},
	)
	if err != nil {
		return nil, err
	}
	var msg = "db migration success"
	return &msg, nil
}

func main() {
	fc.Start(HandleRequest)
}
