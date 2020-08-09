package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//	"github.com/gorilla/mux"
)

/*
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
*/
// Global array to store the different articles
// simulating a database for now
var Articles []Article

func homePage(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles (yup yup yup)")
	json.NewEncoder(writer).Encode(Articles)
}

func main() {
	Articles = []Article{
		{Title: "Hello", Desc: "Art Desc", Content: "My Content"},
		{Title: "Hello", Desc: "Art Desc", Content: "My Content"},
	}
	handleRequests()
}
