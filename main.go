package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	myRouter.HandleFunc("/flights", forwardOpenSkyState)

	// NOTE: Ordering is important here! This has to be defined before
	// the other `/article` endpoint.
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(writer).Encode(Articles)
}

type AirplaneStates struct {
	Time int64 `json:"time"`
	//States []FlightState `json:"states"`
	States [][]interface{} `json:"states"`
}

type Vector struct {
	Time int64
	Name string
}

type MyData struct {
	Dog []Vector
}

func forwardOpenSkyState(writer http.ResponseWriter, req *http.Request) {
	resp, _ := http.Get("https://opensky-network.org/api/states/all")
	body, _ := ioutil.ReadAll(resp.Body)
	var data AirplaneStates
	json.Unmarshal([]byte(body), &data)
	fmt.Printf("Results: %v\n", data.Time)
	fmt.Printf("Results: %v\n", data.States[0])
	fs := FlightState{data: data.States[0]}
	fmt.Printf("Flight State: %v\n", *fs.getLong())
	fmt.Printf("Flight State: %v\n", *fs.getLat())
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Looping through to find the article matching the id
	// TODO: Should use map
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func main() {
	fmt.Println("Rest API v2.0  - Mux Router introduced")
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Art Desc", Content: "My Content"},
		{Id: "2", Title: "Hello 2", Desc: "Art Desc", Content: "My Content"},
	}
	handleRequests()
}
