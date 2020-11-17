package main

import (
	"log"
	"net/http"
)

// @title Movie API
// @version 1.0
// @description Esta es una simple API Rest de peliculas
// @host localhost:8070
// @BasePath /
func main() {
	// Configuracion de Rutas
	router := newRouter()
	// Servidor escuchando en puerto 8070
	server := http.ListenAndServe(":8070", router)
	// Error en caso que termine el servidor
	log.Fatal(server)
}
