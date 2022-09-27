package service

import (
	"fmt"

	"github.com/harrypunk/haikou_cloud/mock_data"
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
)

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

	fams := make([]model.Family, 0)
	for fam := range mockFamilies(100, familyNum,
		randGrade(grades),
		randSchool(schools),
	) {
		fams = append(fams, fam)
	}
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

func randGrade(grades []model.Grade) <-chan model.Grade {
	var ch = make(chan model.Grade)
	go func() {
		for i := 0; ; i = (i + 1) % len(grades) {
			ch <- grades[i]
		}
	}()
	return ch
}

func randSchool(schools []model.School) <-chan model.School {
	var ch = make(chan model.School)
	go func() {
		for i := 0; ; i = (i + 1) % len(schools) {
			ch <- schools[i]
		}
	}()
	return ch
}

func mockFamilies(seed int64, count int,
	grade <-chan model.Grade,
	school <-chan model.School) <-chan model.Family {
	var generator = mock_data.NewWithSeed(seed)
	var names = generator.GetFamilyNames(count)
	var phones = generator.RandomPhone(4 * count)
	var familyCh = make(chan model.Family)
	go func() {
		for famNames := range names {
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
			familyCh <- fam
		}
		close(familyCh)
	}()

	return familyCh
}
