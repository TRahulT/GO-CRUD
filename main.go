package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json: "id"`
	Isbn string `json:"isbn`
	title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname`
	Lastname string `json:"lastname`
}
var movies  []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ =json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","appliction/json")
	params :=  mux.Vars(r)
	for index , items := range movies{
		if items.ID == params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie
		    _ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
		
	
}
func main(){
	r :=mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227",title:"movie one", Director: &Director{Firstname: "john",Lastname: "carter"} })
	movies = append(movies, Movie{ID: "2",Isbn: "875483",title:"movie second",Director: &Director{Firstname: "Rahul",Lastname: "gujjar"}})
	movies = append(movies, Movie{ID: "3",Isbn: "432553",title: "movie third",Director:&Director{Firstname: "Raju",Lastname: "photgrapher"}})

	r.HandleFunc("/movies/",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}/",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}/",updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}/",deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}