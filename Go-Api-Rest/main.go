package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
//json struct
type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}
type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some content",
	},
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}
	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}
	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "La tarea con el id %v  a sido eliminar", taskId)

		}
	}
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	taskId,err:=strconv.Atoi(vars["id"])
	var updateTask task
	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}
	reqBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Fprintf(w,"Please enter valid data")
	}
	json.Unmarshal(reqBody,&updateTask)
	for i,task:=range tasks{
		if task.ID==taskId{
			tasks=append(tasks[:i],tasks[i+1:]...)
			updateTask.ID=taskId
			tasks=append(tasks,updateTask)
			fmt.Fprintf(w,"The task with Id %v has been updated succesfully",taskId)
		}
	}
}




func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "wecolme to my apidddsdddsssffffss")

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")

	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")


	log.Fatal(http.ListenAndServe(":3000", router))
}
