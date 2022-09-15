package service

import (
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

/*
func parent(name string, gender uint8, phone string) model.Parent {
	return model.Parent{
		Name:   name,
		Gender: gender,
		Phone:  phone,
	}
}

func student(name string, gender uint8) model.Student {

}

func mockFamilies(seed int64, count int, bufferSize uint) <-chan []model.Family {
	var generator = mock_data.NewWithSeed(seed)
	var names = generator.GetFamilyNames(count)
	var familyCh = make(chan model.Family)
	for famNames := range names {
		var father = famNames[0]
		var mother = famNames[1]
		var child1 = famNames[2]
		var child2 = famNames[3]
	}

	return bufferFamilies(familyCh, bufferSize)
}

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
