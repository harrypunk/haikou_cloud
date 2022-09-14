package service

import (
	"github.com/harrypunk/haikou_cloud/mock_data"
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
)

func AddMockData(db *gorm.DB) error {
	// dbSession := db.Session(&gorm.Session{CreateBatchSize: 10})

	return nil
}

func AddMockSchool(db *gorm.DB) error {
	var result = db.Create(&mockSchools)
	return result.Error
}

var mockSchools = []model.School{
	school("海南中学"),
	school("海口第一中学"),
	school("海口第二中学"),
	school("秀英第三中学"),
	school("慧蔚外国语"),
}

func school(name string) model.School {
	return model.School{
		Name: name,
	}
}

func parent(name string, gender uint8, phone string) model.Parent {
	return model.Parent{
		Name:   name,
		Gender: gender,
		Phone:  phone,
	}
}

func mockFamilies(seed int64, capacity int, bufferSize int) <-chan []model.Family {
	var generator = mock_data.NewWithSeed(12)
	var ch = make(chan []model.Family, capacity)

	return ch
}
