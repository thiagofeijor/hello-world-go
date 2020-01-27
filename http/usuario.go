package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Usuario model
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler function for /json endpoint
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	usuario := Usuario{
		Nome: "Feijó",
		ID:   1,
	}
	b, _ := json.Marshal(usuario)

	switch {
	case r.Method == "GET":
		fmt.Fprintf(w, string(b))
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry :(")
	}
}
