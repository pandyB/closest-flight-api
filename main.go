package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Global array to store the different articles
// simulating a database for now
var Articles []Article

func homePage(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// create a new mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(writer).Encode(Articles)
}

func main() {
	fmt.Println("Rest API v2.0  - Mux Router introduced")
	Articles = []Article{
		{Title: "Hello", Desc: "Art Desc", Content: "My Content"},
		{Title: "Hello 2", Desc: "Art Desc", Content: "My Content"},
	}
	handleRequests()
}
