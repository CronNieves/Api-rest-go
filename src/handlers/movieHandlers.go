package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CronNieves/Api-rest-go/src/models"
	"github.com/CronNieves/Api-rest-go/src/storage/mongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

// Get root of service
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service unavailable")
}

// GetMovieList Get all Movie
func GetMovieList(w http.ResponseWriter, r *http.Request) {
	var result models.Movies
	var err error
	result, err = storage.MovieList()
	if err != nil {
		w.WriteHeader(500)
	} else {
		makeMoviesResponse(w, 200, result)
	}

}

// GetMovie Get Movie by Id
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	movie, err := storage.MovieFindOneById(oid)
	if err != nil {
		w.WriteHeader(500)
	} else {
		makeMovieResponse(w, 200, movie)
	}
}

// MovieAdd Inser new Movie
func AddMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("")
	decoder := json.NewDecoder(r.Body)
	var movieData models.Movie
	error := decoder.Decode(&movieData)
	if error != nil {
		panic(error)
	}
	defer r.Body.Close()

	errorInsert := storage.MovieAdd(movieData)
	if errorInsert != nil {
		w.WriteHeader(500)
	} else {
		makeMovieResponse(w, 200, movieData)
	}
}

// Update movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)
	decoder := json.NewDecoder(r.Body)

	var movieData models.Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	err = storage.UpdateMovie(oid, movieData)

	if err != nil {
		w.WriteHeader(500)
	} else {
		makeMovieResponse(w, 200, movieData)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	err := storage.DeleteMovie(oid)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(movieId)
	}
}

func makeMovieResponse(w http.ResponseWriter, status int, results models.Movie) {
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Content-Disposition", "null")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func makeMoviesResponse(w http.ResponseWriter, status int, results models.Movies) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}
