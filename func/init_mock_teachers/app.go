package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func initMockTeachers(ctx context.Context, event StructEvent) (*string, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	var client = service.NewMockClient(100, db)
	rows, err := client.AddMockTeachers(15)
	if err != nil {
		return nil, err
	}
	var msg = fmt.Sprintf("add mock teachers %v", rows)
	return &msg, nil
}

func main() {
	fc.Start(initMockTeachers)
}
