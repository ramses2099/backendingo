package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramses2099/backendingo/controllers"
)

func main() {

	// use mux framework
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running")
	})
	// route get all posts
	router.HandleFunc("/api/v1/posts", controllers.GetAllPost).Methods("GET")
	// route get all posts
	router.HandleFunc("/api/v1/posts", controllers.AddPost).Methods("POST")

	// listen in port
	const port string = ":3001"
	log.Println("listening in port ", port)
	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatal(err.Error())
	}

}
