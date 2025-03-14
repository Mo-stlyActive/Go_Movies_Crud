package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	_ "math/rand"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "437556", Title: "Cherry Blossoms",
		Director: &Director{Firstname: "Ruth", LastName: "Bourke"}})
	movies = append(movies, Movie{ID: "2", Isbn: "455676", Title: "Superman",
		Director: &Director{Firstname: "Mahmed", LastName: "Ibrahim"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
