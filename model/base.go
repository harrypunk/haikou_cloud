package model

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Gender   int    `gorm:"type:TINYINT UNSIGNED;"`
	Age      int    `gorm:"type:TINYINT UNSIGNED;"`
	Name     string `gorm:"type:varchar(10);"`
	Phone    string `gorm:"type:varchar(20);"`
	FamilyID uint
	GradeID  uint
	SchoolID uint
	Courses  []*Course  `gorm:"many2many:student_course;"`
	Teachers []*Teacher `gorm:"many2many:student_teacher;"`
}

type Parent struct {
	gorm.Model
	Gender      int    `gorm:"type:TINYINT UNSIGNED;"`
	Name        string `gorm:"type:varchar(10);"`
	Phone       string `gorm:"type:varchar(20);"`
	Description string `gorm:"type:varchar(30);"`
	FamilyID    uint
}

type Family struct {
	gorm.Model
	Name        string `gorm:"type:varchar(10);"`
	Description string `gorm:"type:varchar(30);"`
	Students    []Student
	Parents     []Parent
}

type School struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20);"`
	Description string `gorm:"type:varchar(30);"`
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
	Name     string     `gorm:"type:varchar(10);"`
	Students []*Student `gorm:"many2many:student_course;"`
}

type Teacher struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(20);"`
	Description string     `gorm:"type:TEXT;"`
	Students    []*Student `gorm:"many2many:student_teacher;"`
}
