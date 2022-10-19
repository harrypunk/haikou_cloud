package main

import (
	"context"
	"encoding/json"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/api"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func get_students(ctx context.Context, event StructEvent) ([]byte, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	client, err := service.NewSimpleClient(db)
	if err != nil {
		return nil, err
	}
	var api = api.NewStudentApi(*client)
	studs, err := api.GetStudentList(0, 100)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(studs)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func main() {
	fc.Start(get_students)
}
