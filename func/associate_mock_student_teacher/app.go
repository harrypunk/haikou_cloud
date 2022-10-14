package main

import (
	"context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/harrypunk/haikou_cloud/service"
)

type StructEvent struct {
}

func associateMockStudentTeacher(ctx context.Context, event StructEvent) (*string, error) {
	db, err := service.EnvConnection()
	if err != nil {
		return nil, err
	}
	var client = service.NewMockClient(100, db)
	err = client.AssociateStudentTeacher()
	if err != nil {
		return nil, err
	}
	var msg = "associate student teacher ok"
	return &msg, nil
}

func main() {
	fc.Start(associateMockStudentTeacher)
}
