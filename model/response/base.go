package response

type Entity struct {
}

type StudentInfo struct {
	ID     uint
	Name   string
	Age    uint8
	Gender uint8
	Phone  string
}

type StudentDetail struct {
	ID              uint
	Name            string
	Age             uint8
	Gender          uint8
	Phone           string
	Family          *FamilyInfo
	School          *SchoolInfo
	Courses         []CourseInfo
	CurrentStudents []TeacherInfo
	ReferTeachers   []TeacherInfo
}

type FamilyInfo struct {
	ID      uint
	Parents []ParentInfo
}

type ParentInfo struct {
	ID     uint
	Name   string
	Gender uint8
}

type SchoolInfo struct {
	ID   uint
	Name string
}

type CourseInfo struct {
	ID   uint
	Name string
}

type TeacherInfo struct {
	ID   uint
	Name string
}
