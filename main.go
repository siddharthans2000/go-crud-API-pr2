package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *director `json:"*director"`
}

type director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies[] Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")  // Defining that the application is json
	json.NewEncoder(w).Encode(movies)                  // get all the details and encode it in json to see it as output
}

func getMovie(w http.ResponseWriter, r *http.Request){   
	w.Header().Set("Content-Type","application/json")
	param:=mux.Vars(r)                                  // using the mux package we get the data from the http request we send(Here the Movie for movie details)
	for _,data:= range movies{
		if data.ID==param["id"] {                       //checking which data matched with the param id
			json.NewEncoder(w).Encode(data)             // stroring the requested value in w in json using encoder
			break
		}
	}
}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")    
	param:=mux.Vars(r)                                   // to get which movie id to delete
	for index,item:=range movies{
		if item.ID==param["id"]{
			movies=append(movies[:index],movies[index+1:]...)  // this finds the index and removes the data at that index
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)                    // stores the movie details from the request in a movie variable
	movie.ID=strconv.Itoa(rand.Intn(100000000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param:=mux.Vars(r)
	for index,data:=range movies{
		if data.ID==param["id"] {
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=param["id"]
			movies=append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
var movie[] Movie
func main(){
	par:=mux.NewRouter()

	movies=append(movies,Movie{ID:"1",Isbn: "111234",Title:"Die Hard",Director: &director{Firstname:"David",Lastname: "Fincher"}})
	movies=append(movies,Movie{ID:"2",Isbn:"124411",Title:"Fight Clud",Director: &director{Firstname:"David",Lastname: "Fincher"}})
	par.HandleFunc("/movies",getMovies).Methods("GET")
	par.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	par.HandleFunc("/movies",createMovie).Methods("POST")
	par.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	par.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")

	fmt.Printf("Listening to Port 8081\n")
	log.Fatal(http.ListenAndServe(":8081",par))
}