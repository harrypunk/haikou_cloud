package model

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Gender   int
	Age      int
	Name     string
	Phone    string
	FamilyID uint
	GradeID  uint
	SchoolID uint
	Courses  []*Course  `gorm:"many2many:student_course;"`
	Teachers []*Teacher `gorm:"many2many:student_teacher;"`
}

type Parent struct {
	gorm.Model
	Gender   int
	Name     string
	Phone    string
	FamilyID uint
}

type Family struct {
	gorm.Model
	Name     string
	Students []Student
	Parents  []Parent
}

type School struct {
	gorm.Model
	Name        string
	Description string
	Students    []Student
}

type Grade struct {
	gorm.Model
	Num         int
	DisplayText string
	Students    []Student
}

type Course struct {
	gorm.Model
	Name     string
	Students []*Student `gorm:"many2many:student_course;"`
}

type Teacher struct {
	gorm.Model
	Name     string
	Students []*Student `gorm:"many2many:student_teacher;"`
}
