package main

import (
	"context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func migrate(ctx context.Context, event StructEvent) (*string, error) {
	err := service.Migrate()
	if err != nil {
		return nil, err
	}
	var msg = "db migration success"
	return &msg, nil
}

func main() {
	fc.Start(migrate)
}
