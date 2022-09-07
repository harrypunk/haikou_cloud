package service

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

func setLocalDB() {
	os.Setenv("DB_NAME", "haikou_test")
	os.Setenv("DB_USER", "onetwo")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_URL", "localhost")
}

func localDB() (db *gorm.DB, err error) {
	return Connection("onetwo", "", "localhost", "haikou_test")
}

func exampleData() CommonData {
	return NewCommonData(2)
}

func TestMigrate(t *testing.T) {
	setLocalDB()
	err := Migrate()
	if err != nil {
		t.Fail()
	}
}

func TestGrades(t *testing.T) {
	db, err := localDB()
	if err != nil {
		t.Error(err)
	}
	data := exampleData()
	err = AddGrades(db, &data)
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
	data := exampleData()
	err = AddCourses(db, &data)
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
	err = AddMockSchool(db)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
