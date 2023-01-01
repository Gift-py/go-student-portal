package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gift-py/go-student-portal/pkg/models"
	"github.com/Gift-py/go-student-portal/pkg/utils"
	"github.com/gorilla/mux"
)

// student basic data
func GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := models.GetStudents()
	if err != nil {
		fmt.Println("ERR 1")
		panic(err)
	}
	res, err := json.Marshal(students)
	if err != nil {
		fmt.Println("ERR 2")
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	newStudent := &models.Student{}
	utils.ParseBody(r, newStudent)
	nS, err := models.CreateStudent(newStudent)
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(nS)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	student, err := models.GetStudent(params["sid"])
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	student, err := models.GetStudent(params["sid"])
	if err != nil {
		panic(err)
	}

	var update = &models.Student{}
	utils.ParseBody(r, update)

	if update.Name != "" {
		student.Name = update.Name
	}
	if update.StudentID != "" {
		student.StudentID = update.StudentID
	}
	if update.Age != student.Age {
		student.Age = update.Age
	}
	if update.Level != student.Age {
		student.Level = update.Level
	}

	models.DB.Save(&student)
	res, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	student, err := models.DeleteStudent(params["sid"])
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// course data
func GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := models.GetCourses()
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(courses)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	newCourse := &models.Course{}
	utils.ParseBody(r, newCourse)
	nC, err := models.CreateCourse(newCourse)
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(nC)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	course, err := models.GetCourse(params["cid"])
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	course, err := models.GetCourse(params["cid"])
	if err != nil {
		panic(err)
	}

	update := &models.Course{}
	utils.ParseBody(r, update)

	if update.CourseCode != "" {
		course.CourseCode = update.CourseCode
	}
	if update.CourseName != "" {
		course.CourseName = update.CourseName
	}

	models.DB.Save(&course)
	res, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	course, err := models.DeleteCourse(params["cid"])
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// student courses
func GetStudentCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courses, err := models.GetStudentCourses(params["sid"])
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func AddStudentCourse(w http.ResponseWriter, r *http.Request) {
	type cid_struct struct {
		CourseCodes string `json:"course_code"`
	}
	params := mux.Vars(r)
	courseID := &cid_struct{}
	utils.ParseBody(r, courseID)
	student, err := models.AddStudentCourse(params["sid"], courseID.CourseCodes)
	if err != nil {
		panic(err)
	}
	fmt.Println(student)
	res, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteStudentCourse(w http.ResponseWriter, r *http.Request) {

}
