package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie Struct (Model)
type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Year     string    `json:"year"`
	Director *Director `json:"director"`
}

// Director Struct
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init movie var as a slice Movie struct
var movies []Movie

// Get All Movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "allpication/json")
	json.NewEncoder(w).Encode(movies)
}

// Get Single Movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "allpication/json")
	params := mux.Vars(r) // Get Params
	// loop throngh movies and find with id
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// Add New Movie
func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "allpication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000)) // Mock Id - not safe
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(&Movie{})
}

// Update Movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "allpication/json")
	params := mux.Vars(r)
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
	json.NewEncoder(w).Encode(movies)
}

// Delete Movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "allpication/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implpement DB
	movies = append(movies, Movie{ID: "1", Title: "Avatar", Year: "2009", Director: &Director{FirstName: "James", LastName: "Cameron"}})
	movies = append(movies, Movie{ID: "2", Title: "Iron Man", Year: "2008", Director: &Director{FirstName: "Jon", LastName: "Favreau"}})
	movies = append(movies, Movie{ID: "3", Title: "Thor: Ragnarok", Year: "2017", Director: &Director{FirstName: "Taika", LastName: "Waititi"}})
	movies = append(movies, Movie{ID: "4", Title: "Titanic", Year: "1997", Director: &Director{FirstName: "James", LastName: "Cameron"}})

	// Route Handlers
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", addMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
