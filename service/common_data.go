package service

import (
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddGrades(db *gorm.DB, data *CommonData) error {
	var result = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&data.Grades)
	return result.Error
}

func AddCourses(db *gorm.DB, data *CommonData) error {
	var result = db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&data.Courses)
	return result.Error
}

type CommonData struct {
	Grades  []model.Grade
	Courses []model.Course
}

func NewCommonData(gradeKeyMin int) CommonData {
	var texts = defaultGrades()
	var len1 = len(texts)
	var grades = make([]model.Grade, len1)
	for i := 0; i < len1; i++ {
		grades[i] = grade(uint8(gradeKeyMin+i), texts[i])
	}

	var crs = defaultCourses()
	var len2 = len(crs)
	var courses = make([]model.Course, len2)
	for i := 0; i < len2; i++ {
		courses[i] = course(crs[i])
	}

	return CommonData{
		Grades:  grades,
		Courses: courses,
	}
}

func defaultGrades() []string {
	return []string{"初一", "初二", "初三", "高一", "高二", "高三"}
}

func defaultCourses() []string {
	return []string{"物理", "化学", "数学", "英语"}
}

func course(name string) model.Course {
	return model.Course{
		Name: name,
	}
}

func grade(num uint8, text string) model.Grade {
	return model.Grade{
		Num:         num,
		DisplayText: text,
	}
}
