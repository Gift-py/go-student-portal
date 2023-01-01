package models

import (
	"fmt"

	"github.com/Gift-py/go-student-portal/pkg/config"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// student model
type Student struct {
	gorm.Model
	StudentID string    `json:"sid" gorm:"unique"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Level     int       `json:"level"`
	Courses   []*Course `json:"Courses" gorm:"many2many:student_courses;"`
}

// course model
type Course struct {
	gorm.Model
	CourseCode string `json:"course_code" gorm:"unique"`
	CourseName string `json:"course_name"`
}

type StudentCourse struct {
	student_id uint
	course_id  uint
}

func init() {
	DB = config.Conn()
	DB.AutoMigrate(&Student{})
	DB.AutoMigrate(&Course{})
}

// create student
func CreateStudent(s *Student) (*Student, error) {
	DB.NewRecord(s)
	err := DB.Model(&Student{}).Create(&s).Error
	return s, err
}

// get students
func GetStudents() ([]Student, error) {
	var students []Student
	err := DB.Model(&Student{}).Find(&students).Error
	return students, err
}

// get student
func GetStudent(id string) (*Student, error) {
	var student Student
	err := DB.Model(&Student{}).Where("student_id=?", id).Find(&student).Error
	return &student, err
}

// delete student
func DeleteStudent(id string) (Student, error) {
	var student Student
	err := DB.Model(&Student{}).Where("student_id=?", id).Delete(&student).Error
	return student, err
}

// create course
func CreateCourse(c *Course) (*Course, error) {
	DB.NewRecord(c)
	err := DB.Model(&Course{}).Create(&c).Error
	return c, err
}

// get courses
func GetCourses() ([]Course, error) {
	var courses []Course
	err := DB.Model(&Course{}).Find(&courses).Error
	return courses, err
}

// get course
func GetCourse(id string) (*Course, error) {
	var course Course
	err := DB.Model(&Course{}).Where("course_code=?", id).Find(&course).Error
	return &course, err
}

// delete course
func DeleteCourse(id string) (Course, error) {
	var course Course
	err := DB.Model(&Course{}).Where("course_code=?", id).Delete(&course).Error
	return course, err
}

// student courses
func GetStudentCourses(id string) ([]uint, error) {
	student, err := GetStudent(id)
	if err != nil {
		panic(err)
	}

	var course_ids []uint
	err = DB.Table("student_courses").Select("course_id").Where("student_id=?", student.ID).Find(&course_ids).Error

	return course_ids, err
}

func AddStudentCourse(sid string, cid string) (StudentCourse, error) {
	student, err := GetStudent(sid)
	if err != nil {
		panic(err)
	}
	course, err := GetCourse(cid)
	if err != nil {
		panic(err)
	}

	data := StudentCourse{
		student_id: student.ID,
		course_id:  course.ID,
	}
	fmt.Println(data)
	DB.Table("student_courses").Save(&data)

	return data, err
}

func DeleteStudentCourse(sid string, cid string) (StudentCourse, error) {
	student, err := GetStudent(sid)
	if err != nil {
		panic(err)
	}
	course, err := GetCourse(cid)
	if err != nil {
		panic(err)
	}

	var record StudentCourse

	fin_err := DB.Table("student_courses").Where(map[string]interface{}{"mat_no": student.StudentID, "course_code": course.CourseCode}).Delete(&record).Error

	return record, fin_err
}
