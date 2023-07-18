package routes

import (
	"github.com/fajardodiaz/infosgroup-employee-management/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	R := mux.NewRouter()

	gender := R.PathPrefix("/gender").Subrouter()
	gender.HandleFunc("/", controllers.GetGendersHandler).Methods("GET")
	gender.HandleFunc("/{id:[0-9]+}", controllers.GetGenderHandler).Methods("GET")
	gender.HandleFunc("/", controllers.PostGenderHandler).Methods("POST")
	gender.HandleFunc("/{id}", controllers.PutGenderHandler).Methods("PUT")
	gender.HandleFunc("/{id}", controllers.DeleteGenderHandler).Methods("DELETE")

	return R
}
