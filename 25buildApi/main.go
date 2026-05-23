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

// model for course -file

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

//fake DB

// courses :=[]course
var courses []Course

// MiddleWare,helper-file

func (c *Course) IsEmpty() bool {

	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("WELCOME TO GO API")
	r := mux.NewRouter()

	//SEEDING
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 288, Author: &Author{Fullname: "Ahmed Abdilahi", Website: "www.google.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "GO Lang", CoursePrice: 210, Author: &Author{Fullname: "Abdullahi Ahmed", Website: "www.telesom.net"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "JAVASCRIPTS", CoursePrice: 100, Author: &Author{Fullname: "Abdirahman Ahmed", Website: "www.youtube.com"}})

	// routing

	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course {id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOnecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updatecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOnecourse).Methods("DELETE")
	// listen port

	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers --file

// serve home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API in GO lang"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// function

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get ONE COURSE")
	w.Header().Set("Content-Type", "application/json")

	// GRAB ID FROM REQUEST

	params := mux.Vars(r)

	// loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with geiven id")
	return
}

//create one course

func createOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created course one")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about -{}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// generate unique ID ,string

	//append course into courses

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

//update course

func updatecourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("created course one")
	w.Header().Set("Content-Type", "application/json")

	//first -grab if from req

	params := mux.Vars(r)

	// loop , id ,remove the id,add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	/// TODO send a response when id is not found

}

//delete course

func deleteOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// loop ,id,remove,index,index1

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
