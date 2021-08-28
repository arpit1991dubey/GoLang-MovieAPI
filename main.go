package main

import (
	"fmt"
	"encoding/json" 		 //To encode and decode json
	"log"					 //To log Errors
	"github.com/gorilla/mux" //The name mux stands for "HTTP request multiplexer".
				//Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes 
				//and calls a handler for the route that matches the URL or other conditions
	"math/rand"              // To generate random values
	"net/http"               //To create an http server
	"strconv"                //To convert string into integrs or vice verso
							

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

/******************Function to get the list of all movies**************************/

func getMovies(w http.ResponseWriter, r *http.Request) { //W is the response and r is the request
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

/******************Function to delete a single movie**************************/

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

/******************Function to get a single movie**************************/

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

/******************Function to create a new movie**************************/

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

/******************Function to update a new movie**************************/

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set the json content type
	w.Header().Set("Content-type", "application/json")
	//get params
	params := mux.Vars(r)

	//loop over the movies,range
	//delete the movie with the meovie id received
	//add a new movie with the body that we receive from the frontend
	//this process of updation is not recommneed for database

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

/******************Main Function to aggregate methods and functions**************************/
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "19145", Title: "Movie two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("STARTING SERVER AT PORT 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
