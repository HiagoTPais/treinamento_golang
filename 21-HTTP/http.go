package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola Mundo"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pagina de usuarios"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
