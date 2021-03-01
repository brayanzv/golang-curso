package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"sin limites", 2012, "desconocido"},
	Movie{"duro de matar", 2000, "Juan Antonio"},
	Movie{"el destructor del mar", 2006, "Antonio"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde el index")
}

func MovieList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
	//devolver en json
	//fmt.Fprintf(w, "bienvenido a la lista de peliculas que tenemos para ofrecerte")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "disfruta de la pelicula que elegiste %s", movie_id)

}
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)

	movies = append(movies, movie_data)
}
