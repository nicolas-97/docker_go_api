package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Manejador para la ruta raíz ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola, mundo!")
	})

	// Inicia el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
