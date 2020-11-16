package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CronNieves/Api-rest-go/src/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return session
}

var movies = models.Movies{
	models.Movie{"IronMan", 2013, "Desconocido"},
	models.Movie{"Batman", 2005, "Pedro Antonio"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service unavailable")
}

func GetMovieList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Obteniendo pelicula con id "+movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("")
	decoder := json.NewDecoder(r.Body)
	var movie_data models.Movie
	error := decoder.Decode(&movie_data)
	if error != nil {
		panic(error)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)

}
