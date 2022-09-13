package main

import (
	"context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func addCommon(ctx context.Context, event StructEvent) (*string, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	data := service.NewCommonData()
	err = service.AddCourses(db, &data)
	if err != nil {
		return nil, err
	}
	err = service.AddGrades(db, &data)
	if err != nil {
		return nil, err
	}
	var msg = "grades and courses added"
	return &msg, nil
}

func main() {
	fc.Start(addCommon)
}
