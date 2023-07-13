package main

import (
	"log"
	"net/http"

	"github.com/fajardodiaz/infosgroup-employee-management/controllers"
	"github.com/fajardodiaz/infosgroup-employee-management/initializer"
	"github.com/gorilla/mux"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
}

func main() {

	r := mux.NewRouter()

	gender := r.PathPrefix("/gender").Subrouter()
	gender.HandleFunc("/", controllers.GetGendersHandler).Methods("GET")
	gender.HandleFunc("/{id:[0-9]+}", controllers.GetGenderHandler).Methods("GET")
	gender.HandleFunc("/", controllers.PostGenderHandler).Methods("POST")
	gender.HandleFunc("/{id}", controllers.PutGenderHandler).Methods("PUT")
	gender.HandleFunc("/{id}", controllers.DeleteGenderHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
