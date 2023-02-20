package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// capital struct name bhelps to export it if it is in another file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"aurthor"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middleware for empty field check
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}
func main() {
	fmt.Println("API try in GoLang")
	r := mux.NewRouter()

	//Fillinf Pseudo data
	courses = append(courses, Course{
		CourseId: "1", CourseName: "ReactJs", CoursePrice: 200, Author: &Author{
			Fullname: "Naman",
			Website:  "hello.co.in",
		},
	})

	courses = append(courses, Course{
		CourseId: "2", CourseName: "GoLang", CoursePrice: 300, Author: &Author{
			Fullname: "Hitesh",
			Website:  "Hi.co.in",
		},
	})

	//routing
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/createCourse", createOneCourse).Methods("POST")
	r.HandleFunc("/update/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteOneCourse).Methods("DELETE")

	//establishing Connection
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection established succesfully")
	}
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}
func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses..")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id fron request
	params := mux.Vars(r)

	for i := range courses {
		if courses[i].CourseId == params["id"] {
			json.NewEncoder(w).Encode(courses[i])
			return
		}
	}

	json.NewEncoder(w).Encode("No course found based on given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please add course name..")
		return
	}

	//generate unique id, string
	//append course int courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

// check it once again
func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	//In this we remove old entry and add the new updated entry
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from the request

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			//remove the course
			courses = append(courses[:i], courses[i+1:]...)

			//adding new course
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, course := range courses {
		if course.CourseId == params["id"] {
			//delete the course
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
}
