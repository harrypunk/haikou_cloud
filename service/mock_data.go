package service

import (
	"fmt"

	"github.com/harrypunk/haikou_cloud/mock_data"
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
)

func AddMockData(db *gorm.DB, familyNum int) error {
	var fams = make([]model.Family, familyNum)
	for fam := range mockFamilies(100, familyNum) {
		fams = append(fams, fam)
	}
	var result = db.CreateInBatches(fams, 10)

	return result.Error
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

func parent(name string, gender uint8, phone string) model.Parent {
	return model.Parent{
		Name:   name,
		Gender: gender,
		Phone:  phone,
	}
}

func student(name string, gender uint8, phone string) model.Student {
	return model.Student{
		Name:   name,
		Gender: gender,
		Phone:  phone,
	}
}

func mockFamilies(seed int64, count int) <-chan model.Family {
	var generator = mock_data.NewWithSeed(seed)
	var names = generator.GetFamilyNames(count)
	var phones = generator.RandomPhone(4 * count)
	var familyCh = make(chan model.Family)
	for famNames := range names {
		var father = famNames[0]
		var mother = famNames[1]
		var child1 = famNames[2]
		var child2 = famNames[3]

		var fam = model.Family{
			Name: fmt.Sprintf("家庭 %v,%v", father, mother),
			Students: []model.Student{
				student(child1, 1, <-phones),
				student(child2, 2, <-phones),
			},
			Parents: []model.Parent{
				parent(father, 1, <-phones),
				parent(mother, 2, <-phones),
			},
		}
		familyCh <- fam
	}
	close(familyCh)

	return familyCh
}

/*
func bufferFamilies(ch chan model.Family, bufferSize uint) <-chan []model.Family {
	var outCh = make(chan []model.Family)
	buffer := make([]model.Family, bufferSize)
	for fam := range ch {
		buffer = append(buffer, fam)
		if len(buffer) >= int(bufferSize) {
			outCh <- buffer
			buffer = make([]model.Family, bufferSize)
		}
	}
	if len(buffer) > 0 {
		outCh <- buffer
	}
	close(outCh)
	return outCh
}
*/
