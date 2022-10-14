package service

import (
	"fmt"

	"math/rand"

	"github.com/harrypunk/haikou_cloud/mock_data"
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
)

type MockClient struct {
	rand rand.Rand
	db   *gorm.DB
}

func NewMockClient(seed int64, db *gorm.DB) MockClient {
	return MockClient{
		rand: *rand.New(rand.NewSource(seed)),
		db:   db,
	}
}

func (client *MockClient) AssociateStudentTeacher() error {
	// all teachers
	var teachers []model.Teacher
	result := client.db.Find(&teachers)
	err := result.Error
	if err != nil {
		return err
	}
	// all students
	var students []model.Student
	result = client.db.Find(&students)
	err = result.Error
	if err != nil {
		return err
	}

	var teacherCh = loopArray(teachers)
	for _, st := range students {
		var currentTeacherCount = client.rand.Intn(2) + 1
		var referTeacherCount = client.rand.Intn(3) + 1
		var currentTeacherList []model.Teacher
		for i := 0; i < currentTeacherCount; i++ {
			teacher := <-teacherCh
			currentTeacherList = append(currentTeacherList, teacher)
		}
		client.db.Model(&st).Association("CurrentTeachers").Append(currentTeacherList)

		var referTeacherList []model.Teacher
		for i := 0; i < referTeacherCount; i++ {
			teacher := <-teacherCh
			referTeacherList = append(referTeacherList, teacher)
		}
		client.db.Model(&st).Association("ReferTeachers").Append(referTeacherList)
	}
	return nil
}

func (client *MockClient) AddMockTeachers(count int) (*int64, error) {
	var courses []model.Course
	result := client.db.Find(&courses)
	err := result.Error
	if err != nil {
		return nil, err
	}
	var teacherNames = mockTeacherNames(103)
	var courseCh = loopArray(courses)
	teachers := make([]model.Teacher, 0)
	for i := 0; i < count; i++ {
		teacher := <-teacherNames
		teacher.CourseId = (<-courseCh).ID
		teachers = append(teachers, teacher)
	}
	result = client.db.Create(teachers)
	return &result.RowsAffected, result.Error
}

func (client *MockClient) AddMockSessions(count int) (*int64, error) {
	// get all teachers
	var teachers []model.Teacher
	result := client.db.Find(&teachers)
	err := result.Error
	if err != nil {
		return nil, err
	}
	// get all students
	var students []model.Student
	result = client.db.Find(&students)
	err = result.Error
	if err != nil {
		return nil, err
	}
	var teacherCh = loopArray(teachers)
	var studCh = loopArray(students)
	sessions := make([]model.Session, count)
	for i := 0; i < count; i++ {
		var s = model.Session{}
		mainTeacher := <-teacherCh
		s.CourseId = mainTeacher.CourseId
		s.MainTeacherID = mainTeacher.ID
		// random students
		var studentCount = client.rand.Intn(4)
		for j := 0; j < studentCount; j++ {
			st1 := <-studCh
			s.Students = append(s.Students, &st1)
		}
		sessions[i] = s
	}
	result = client.db.Create(sessions)
	return &result.RowsAffected, result.Error
}

func (client *MockClient) AddMockFamilies(familyNum int) error {
	var db = client.db
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

func (client *MockClient) AddMockSchool() error {
	var result = client.db.Create(&mockSchools)
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
