package routes

import (
	"github.com/Gift-py/go-student-portal/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(route *mux.Router) {
	//student basic data
	route.HandleFunc("/students/", controllers.GetStudents).Methods("GET")
	route.HandleFunc("/students/", controllers.CreateStudent).Methods("POST")
	route.HandleFunc("/students/{sid}", controllers.GetStudent).Methods("GET")
	route.HandleFunc("/students/{sid}", controllers.UpdateStudent).Methods("PUT")
	route.HandleFunc("/students/{sid}", controllers.DeleteStudent).Methods("DELETE")

	//student courses
	route.HandleFunc("student/courses/{sid}", controllers.GetStudentCourses).Methods("GET")
	route.HandleFunc("/students/courses/{sid}", controllers.AddStudentCourse).Methods("PUT")
	route.HandleFunc("/Students/courses/{sid}", controllers.DeleteStudentCourse).Methods("DELETE")

	//school courses
	route.HandleFunc("/courses/", controllers.GetCourses).Methods("GET")
	route.HandleFunc("/courses/", controllers.CreateCourse).Methods("POST")
	route.HandleFunc("/courses/{cid}", controllers.GetCourse).Methods("GET")
	route.HandleFunc("/courses/{cid}", controllers.UpdateCourse).Methods("PUT")
	route.HandleFunc("/courses/{cid}", controllers.DeleteCourse).Methods("DELETE")

}
