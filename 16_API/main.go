package main

import (
	"encoding/json"
	"net/http"
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
	return c.CourseId == "" && c.CourseName == "";
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