package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Auther *Auther `json:"auther"`
}

type Auther struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course;

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return  c.CourseName == "";
}

func main() {

}

// controllers - file

// serve home router

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Home.</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func GetOneCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	param := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == param["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the id")
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// if r.Body ==  {
	// 	json.NewEncoder(w).Encode("Please send some data")
	// }

	var course Course;

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("JSON is empty")
		return;
	}

	// gen random courseId
	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append course into courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get param value
	// params := mux.Vars(r)
}