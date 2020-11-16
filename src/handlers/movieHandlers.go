package handlers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service unavailable")
}

func GetMovieList(w http.ResponseWriter, r *http.Request) {
	movies := models.Movies{
		models.Movie{"IronMan", 2013, "Desconocido"},
		models.Movie{"Batman", 2005, "Pedro Antonio"},
	}
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Obteniendo pelicula con id "+movieId)
}
