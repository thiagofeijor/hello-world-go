package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/when", func(w http.ResponseWriter, r *http.Request) {
		s := time.Now().Format("02/01/2006 03:04:05")
		fmt.Fprintf(w, "Hora: %s", s)
	})

	http.HandleFunc("/json", UsuarioHandler)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
