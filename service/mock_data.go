package service

import (
	"fmt"

	"github.com/harrypunk/haikou_cloud/mock_data"
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
)

func AddMockTeachers(db *gorm.DB, count int) *gorm.DB {
	var courses []model.Course
	result := db.Find(&courses)
	err := result.Error
	if err != nil {
		return result
	}
	var teacherNames = mockTeacherNames(103)
	var courseCh = loopArray(courses)
	teachers := make([]model.Teacher, 0)
	for i := 0; i < count; i++ {
		teacher := <-teacherNames
		teacher.CourseId = (<-courseCh).ID
		teachers = append(teachers, teacher)
	}
	return db.Create(teachers)
}

func AddMockData(db *gorm.DB, familyNum int) error {
	var grades []model.Grade
	result := db.Find(&grades)
	err := result.Error
	if err != nil {
		return err
	}

	var schools []model.School
	result = db.Find(&schools)
	err = result.Error
	if err != nil {
		return err
	}

	var fams = mockFamilies(100, familyNum, loopArray(grades), loopArray(schools))
	result = db.CreateInBatches(fams, 10)
	err = result.Error

	return err
}

func AddMockSchool(db *gorm.DB) error {
	var result = db.Create(&mockSchools)
	return result.Error
}

var mockSchools = []model.School{
	school(33, "海南中学"),
	school(34, "海口第一中学"),
	school(35, "海口第二中学"),
	school(36, "秀英第三中学"),
	school(37, "慧蔚外国语"),
}

func school(id uint, name string) model.School {
	return model.School{
		Model: gorm.Model{ID: id},
		Name:  name,
	}
}

func loopArray[T interface{}](arr []T) <-chan T {
	var ch = make(chan T)
	go func() {
		for i := 0; ; i = (i + 1) % len(arr) {
			ch <- arr[i]
		}
	}()
	return ch
}

func mockFamilies(seed int64, count int,
	grade <-chan model.Grade,
	school <-chan model.School) []model.Family {
	var generator = mock_data.NewWithSeed(seed)
	var names = generator.GetFamilyNames()
	var phones = generator.RandomPhone()
	families := make([]model.Family, 0)
	for i := 0; i < count; i++ {
		var famNames = <-names
		var father = famNames[0]
		var mother = famNames[1]
		var child1 = famNames[2]
		var child2 = famNames[3]

		var fam = model.Family{
			Name: fmt.Sprintf("家庭 %v,%v", father, mother),
			Students: []model.Student{
				{
					Name:     child1,
					Gender:   1,
					Phone:    <-phones,
					GradeID:  (<-grade).ID,
					SchoolID: (<-school).ID,
				},
				{
					Name:     child2,
					Gender:   2,
					Phone:    <-phones,
					GradeID:  (<-grade).ID,
					SchoolID: (<-school).ID,
				},
			},
			Parents: []model.Parent{
				{
					Name:   father,
					Gender: 1,
					Phone:  <-phones,
				},
				{
					Name:   mother,
					Gender: 2,
					Phone:  <-phones,
				},
			},
		}
		families = append(families, fam)
	}

	return families
}

func mockTeacherNames(seed int64) <-chan model.Teacher {
	var ch = make(chan model.Teacher)
	var generator = mock_data.NewWithSeed(seed)
	var names = generator.RandomNameList()
	go func() {
		for {
			ch <- model.Teacher{
				Name: <-names,
			}
		}
	}()
	return ch
}
