package service

import (
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddMockData(db *gorm.DB) error {
	return nil
}

func AddGrades(db *gorm.DB) error {
	var texts = defaultGrades()
	var len1 = len(texts)
	var grades = make([]model.Grade, len1)
	for i := 0; i < len1; i++ {
		grades[i] = grade(uint(100+i), uint8(6+i), texts[i])
	}
	var result = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&grades)
	return result.Error
}

func AddCourses(db *gorm.DB) error {
	var courses = []model.Course{
		course(101, "物理"),
		course(102, "化学"),
		course(103, "数学"),
		course(104, "英语"),
	}
	var result = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&courses)
	return result.Error
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

func grade(id uint, num uint8, text string) model.Grade {
	return model.Grade{
		Model:       gorm.Model{ID: id},
		Num:         num,
		DisplayText: text,
	}
}

func defaultGrades() []string {
	return []string{"初一", "初二", "初三", "高一", "高二", "高三"}
}

func course(id uint, name string) model.Course {
	return model.Course{
		Model: gorm.Model{ID: id},
		Name:  name,
	}

}
