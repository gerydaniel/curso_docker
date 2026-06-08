package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Sesion 4 - Imagen optimizada con Go")
	})

	log.Println("Servidor iniciado en puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
