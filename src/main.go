package main

import (
	"log"
	"net/http"
)

func main() {
	// Configuracion de Rutas
	router := newRouter()
	// Servidor escuchando en puerto 8070
	server := http.ListenAndServe(":8070", router)
	// Error en caso que termine el servidor
	log.Fatal(server)
}
