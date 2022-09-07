package service

import (
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddMockData(db *gorm.DB) error {
	dbSession := db.Session(&gorm.Session{CreateBatchSize: 10})

	return nil
}

func AddMockSchool(db *gorm.DB) error {
	var schools = []model.School{
		school(33, "海南中学"),
		school(34, "海口第一中学"),
		school(35, "海口第二中学"),
		school(36, "秀英第三中学"),
		school(37, "慧蔚外国语"),
	}
	var result = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&schools)
	return result.Error
}

func school(id uint, name string) model.School {
	return model.School{
		Model: gorm.Model{ID: id},
		Name:  name,
	}
}

func parent(name string, gender uint8, phone string) model.Parent {
	return model.Parent{
		Name:   name,
		Gender: gender,
		Phone:  phone,
	}
}
