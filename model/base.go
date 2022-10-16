package model

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Gender          uint8  `gorm:"type:TINYINT UNSIGNED;"`
	Age             uint8  `gorm:"type:TINYINT UNSIGNED;"`
	Name            string `gorm:"type:varchar(10);"`
	Phone           string `gorm:"type:varchar(20);"`
	FamilyID        uint
	GradeID         uint
	SchoolID        uint
	School          School
	Courses         []*Course  `gorm:"many2many:student_course;"`
	Sessions        []*Session `gorm:"many2many:student_session;"`
	CurrentTeachers []*Teacher `gorm:"many2many:student_current_teacher"`
	ReferTeachers   []*Teacher `gorm:"many2many:student_refer_teacher"`
}

type Parent struct {
	gorm.Model
	Gender      uint8  `gorm:"type:TINYINT UNSIGNED;"`
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
	Num         uint8  `gorm:"type:TINYINT UNSIGNED;"`
	DisplayText string `gorm:"type:varchar(10);"`
	Students    []Student
}

type Course struct {
	gorm.Model
	Name     string     `gorm:"type:varchar(10);"`
	Students []*Student `gorm:"many2many:student_course;"`
	Teachers []Teacher
	Sessions []Session
}

type Teacher struct {
	gorm.Model
	Name            string `gorm:"type:varchar(20);"`
	Description     string `gorm:"type:TEXT;"`
	CourseId        uint
	MainSessions    []Session  `gorm:"foreignKey:MainTeacherID"`
	CurrentStudents []*Student `gorm:"many2many:student_current_teacher"`
	ReferStudents   []*Student `gorm:"many2many:student_refer_teacher"`
}

type Session struct {
	gorm.Model
	CourseId      uint
	Description   string     `gorm:"type:TEXT;"`
	Students      []*Student `gorm:"many2many:student_session;"`
	MainTeacherID uint
}
