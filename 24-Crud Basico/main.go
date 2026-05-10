package main

import (
	"crud-basico/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Running on 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
