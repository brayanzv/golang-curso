package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", contact)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde el index")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola esta es la pagina de contactos")
}
