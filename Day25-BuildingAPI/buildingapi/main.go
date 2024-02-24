package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Creating Model for Course -File

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// Middleware , helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

//Controllers -go in own seperate Files/Folders
//Serve Home Route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API's</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func main() {

}
