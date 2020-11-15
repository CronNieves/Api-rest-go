package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/movies",getMovieList)
	router.HandleFunc("/movie/{id}",getMovie)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Service unavailable")
	})*/

	server := http.ListenAndServe(":8070",router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Service unavailable")
}

func getMovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Service unavailable")
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params:=mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w,"Obteniendo pelicula con id " + movieId)
}