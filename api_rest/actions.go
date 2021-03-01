package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde el index")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"sin limites", 2012, "desconocido"},
		Movie{"duro de matar", 2000, "Juan Antonio"},
		Movie{"el destructor del mar", 2006, "Antonio"},
	}
	json.NewEncoder(w).Encode(movies)
	//devolver en json
	//fmt.Fprintf(w, "bienvenido a la lista de peliculas que tenemos para ofrecerte")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "disfruta de la pelicula que elegiste %s", movie_id)

}
