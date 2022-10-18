package api

import (
	"github.com/harrypunk/haikou_cloud/model/response"
	"github.com/harrypunk/haikou_cloud/service"
)

type StudentApi struct {
	client service.Client
}

func NewStudentApi(cl service.Client) *StudentApi {
	return &StudentApi{
		client: cl,
	}
}

func (api *StudentApi) GetStudentList(offset int, limit int) ([]response.StudentInfo, error) {
	studs, err := api.client.GetStudents(offset, limit)
	if err != nil {
		return nil, err
	}
	return response.StudentsToInfos(studs), nil
}
