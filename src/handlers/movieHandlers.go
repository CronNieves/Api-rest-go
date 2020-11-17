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

// GetMovieList godoc
// @Description Obtiene todas las peliculas
// @Tags Movie
// @Produce  json
// @Success 200 {object} models.Movies
// @Router /movies [get]
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

// GetMovie godoc
// @Description Obtiene una pelicula por id
// @Tags Movie
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Movie
// @Param id path string true "Movie ID"
// @Router /movie/{id} [get]
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

// AddMovie godoc
// @Description Obtiene una pelicula por id
// @Tags Movie
// @Accept  json
// @Produce  json
// @Param Movie body models.Movie true "New Movie"
// @Success 200 {object} models.Movie
// @Router /movie [POST]
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

// UpdateMovie godoc
// @Description Actualiza una pelicula
// @Tags Movie
// @Accept  json
// @Produce  json
// @Param id path string true "Movie ID"
// @Param Movie body models.Movie true "Update Movie"
// @Success 200 {object} models.Movie
// @Router /movie/{id} [PUT]
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

// DeleteMovie godoc
// @Description Borra una pelicula por id
// @Tags Movie
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Param id path string true "Movie ID"
// @Router /movie/{id} [delete]
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
