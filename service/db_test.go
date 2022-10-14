package service_test

import (
	"os"
	"testing"

	s "github.com/harrypunk/haikou_cloud/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setLocalDB() {
	os.Setenv("DB_NAME", "haikou_test")
	os.Setenv("DB_USER", "onetwo")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_URL", "localhost")
}

func localDB() (db *gorm.DB, err error) {
	return localDBConfig.ToClient(&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

var localDBConfig = s.DBConfig{
	Username: "onetwo",
	Password: "",
	Url:      "localhost",
	DBName:   "haikou_test",
}

func TestMigrate(t *testing.T) {
	setLocalDB()
	err := s.Migrate()
	if err != nil {
		t.Fail()
	}
}

func TestGrades(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	data := s.NewCommonData()
	err = s.AddGrades(db, &data)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCourses(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	data := s.NewCommonData()
	err = s.AddCourses(db, &data)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSchools(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	var client = s.NewMockClient(100, db)
	err = client.AddMockSchool()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestParentsChildren(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	var client = s.NewMockClient(100, db)
	err = client.AddMockFamilies(50)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestTeachers(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	var client = s.NewMockClient(100, db)
	rows, err := client.AddMockTeachers(10)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(rows)
}

func TestSessions(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	var client = s.NewMockClient(100, db)
	rows, err := client.AddMockSessions(12)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(rows)
}

func TestAssociateStudentTeacher(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	var client = s.NewMockClient(100, db)
	err = client.AssociateStudentTeacher()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetStudentList(t *testing.T) {
	client, err := s.NewClient(true, localDBConfig)
	if err != nil {
		t.Error(err)
	}
	students, err := client.GetStudents(1, 10)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("students length: %v", len(students))
}
