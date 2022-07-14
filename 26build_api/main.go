package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Course Model for course - file
type Course struct {
	Id          int     `json:"id"`
	CourseName  string  `json:"course_name"`
	CoursePrice float64 `json:"course_price"`
	Author      *Author `json:"author"`
}

// Fake data for course
var courses = []Course{
	{Id: 1, CourseName: "Go", CoursePrice: 100.0, Author: &authors[0]},
	{Id: 2, CourseName: "Python", CoursePrice: 200.0, Author: &authors[1]},
	{Id: 3, CourseName: "Ruby", CoursePrice: 300.0, Author: &authors[2]},
	{Id: 4, CourseName: "Java", CoursePrice: 400.0, Author: &authors[3]},
	{Id: 5, CourseName: "C++", CoursePrice: 500.0, Author: &authors[4]},
	{Id: 6, CourseName: "C#", CoursePrice: 600.0, Author: &authors[5]},
	{Id: 7, CourseName: "JavaScript", CoursePrice: 700.0, Author: &authors[6]},
	{Id: 8, CourseName: "PHP", CoursePrice: 800.0, Author: &authors[7]},
	{Id: 9, CourseName: "Swift", CoursePrice: 900.0, Author: &authors[8]},
	{Id: 10, CourseName: "Objective-C", CoursePrice: 1000.0, Author: &authors[9]},
}

// Fake data for author
var authors = []Author{
	{Id: 1, Name: "John", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174825?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 2, Name: "Jane", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174826?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 3, Name: "Jack", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174827?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 4, Name: "Joe", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174828?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 5, Name: "Ratul", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174829?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 6, Name: "Hasan", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174830?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 7, Name: "Arroby", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174831?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 8, Name: "Tuly", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174832?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 9, Name: "Raj", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174833?v=3&s=460", Website: "ratulhasan.com"},
	{Id: 10, Name: "Raju", Email: "jhon@doe.com", Avatar: "https://avatars0.githubusercontent.com/u/174834?v=3&s=460", Website: "ratulhasan.com"},
}

// IsCourseEmpty Middleware for course
func (c Course) IsCourseEmpty() bool {
	return c.CourseName == ""
}

//IsAuthorEmpty Middleware for author
func (a Author) IsAuthorEmpty() bool {
	return a.Name == ""
}

type Author struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Avatar  string `json:"avatar"`
	Website string `json:"website"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeController).Methods("GET")
	r.HandleFunc("/courses", GetCoursesController).Methods("GET")
	r.HandleFunc("/course/{id}", GetACourse).Methods("GET")
	r.HandleFunc("/course", CreateCourseController).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourseController).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourseController).Methods("DELETE")

	// Serve the server on port 2000
	// Listen to port 2000
	log.Fatal(http.ListenAndServe(":2000", r))
}

// HomeController for course
func HomeController(w http.ResponseWriter, r *http.Request) {
	write, err := w.Write([]byte("<h1>Welcome to the Home Page</h1>"))
	if err != nil {
		return
	}

	fmt.Println(write)
}

// GetCoursesController for course
func GetCoursesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(courses)
	if err != nil {
		return
	}

	fmt.Println("GetAllCourses")
}

func GetACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)
	for _, course := range courses {
		// convert string to int
		courseId, _ := strconv.Atoi(params["id"])
		if course.Id == courseId {
			err := json.NewEncoder(w).Encode(course)
			if err != nil {
				return
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode("Course not found")
	if err != nil {
		return
	}

	fmt.Println("GetACourse")
}

// CreateCourseController for course
func CreateCourseController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateCourseController")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// What if: r.body is empty
	if r.Body == nil {
		err := json.NewEncoder(w).Encode("Please send a request body")
		if err != nil {
			return
		}
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		return
	}

	// create a course id
	course.Id = len(courses) + 1
	if course.IsCourseEmpty() {
		_ = json.NewEncoder(w).Encode("Course name is required")
		return
	}

	// create a author id
	course.Author.Id = len(authors) + 1

	// check if author is empty
	if course.Author.IsAuthorEmpty() {
		_ = json.NewEncoder(w).Encode("Author name is required")
		return
	}

	courses = append(courses, course)
	err = json.NewEncoder(w).Encode(courses)
	if err != nil {
		return
	}
}

// UpdateCourseController for course
func UpdateCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// What if: r.body is empty
	if r.Body == nil {
		err := json.NewEncoder(w).Encode("Please send a request body")
		if err != nil {
			return
		}
	}

	params := mux.Vars(r)
	for index, course := range courses {
		// convert string to int
		courseId, _ := strconv.Atoi(params["id"])
		if course.Id == courseId {
			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil {
				return
			}

			// remove the course from the array
			courses = append(courses[:index], courses[index+1:]...)
			course.Id = courseId
			// add the course to the array
			courses = append(courses, course)
			err = json.NewEncoder(w).Encode(courses)
			if err != nil {
				return
			}
			return
		}
	}

	// if courses not found
	err := json.NewEncoder(w).Encode("Course not found")
	if err != nil {
		return
	}

	fmt.Println("UpdateCourse")
}

// DeleteCourseController for course
func DeleteCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)
	for index, course := range courses {
		// convert string to int
		courseId, _ := strconv.Atoi(params["id"])
		if course.Id == courseId {
			// remove the course from the array
			courses = append(courses[:index], courses[index+1:]...)
			err := json.NewEncoder(w).Encode(courses)
			if err != nil {
				return
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode("Course not found")
	if err != nil {
		return
	}

	fmt.Println("DeleteCourse")
}
