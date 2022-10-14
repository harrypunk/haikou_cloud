package main

import (
	"context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func addMockData(ctx context.Context, event StructEvent) (*string, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	var client = service.NewMockClient(100, db)
	client.AddMockFamilies(100)
	if err != nil {
		return nil, err
	}
	var msg = "parent and children added"
	return &msg, nil
}

func main() {
	fc.Start(addMockData)
}
