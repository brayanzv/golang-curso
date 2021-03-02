package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSesion() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return session
}

var collection = getSesion().DB("curso-go").C("movies")

var movies = Movies{
	Movie{"sin limites", 2012, "desconocido"},
	Movie{"duro de matar", 2000, "Juan Antonio"},
	Movie{"el destructor del mar", 2006, "Antonio"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde el index")
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("resultados ", results)
	}

	responseMovies(w, 200, results)
	//devolver en json
	//fmt.Fprintf(w, "bienvenido a la lista de peliculas que tenemos para ofrecerte")
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}

	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	old := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(old).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	responseMovie(w, 200, results)

}
func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	old := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(old)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	results := Message{"success", "la pelicula con " + movie_id + " ha sido borrada correctamente"}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

	//responseMovie(w, 200, results.Message())
}
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movie_data)

}

func responseMovie(w http.ResponseWriter, status int, results Movie) {

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)

}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)

}
