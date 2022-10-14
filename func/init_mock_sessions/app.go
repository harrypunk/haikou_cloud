package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func initMockSessions(ctx context.Context, event StructEvent) (*string, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	var client = service.NewMockClient(100, db)
	rows, err := client.AddMockSessions(15)
	if err != nil {
		return nil, err
	}
	var msg = fmt.Sprintf("init mock sessions success %v", rows)
	return &msg, nil
}

func main() {
	fc.Start(initMockSessions)
}
