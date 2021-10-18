package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type task struct{
	ID int `json:ID`
	Name string `json:Name`
	Content string `json:Content`
}
type allTasks []task

var tasks = allTasks{
	{
		ID:1,
		Name:"Task One",
		Content:"Some content",
	
	},
		
}
func indexRoute(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"wecolme to my api")

}

func main(){
	
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",indexRoute)
	log.Fatal(http.ListenAndServe(":3000",router))
}