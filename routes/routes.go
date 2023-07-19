package routes

import (
	"github.com/fajardodiaz/infosgroup-employee-management/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	R := mux.NewRouter()

	v1 := R.PathPrefix("/api/v1").Subrouter()

	// Gender routes
	gender := v1.PathPrefix("/genders").Subrouter()

	gender.HandleFunc("", controllers.GetGendersHandler).Methods("GET")
	gender.HandleFunc("/{id:[0-9]+}", controllers.GetGenderHandler).Methods("GET")
	gender.HandleFunc("", controllers.PostGenderHandler).Methods("POST")
	gender.HandleFunc("/{id}", controllers.PutGenderHandler).Methods("PUT")
	gender.HandleFunc("/{id}", controllers.DeleteGenderHandler).Methods("DELETE")

	// Position routes
	position := v1.PathPrefix("/positions").Subrouter()
	position.HandleFunc("", controllers.GetPositionsHandler).Methods("GET")
	position.HandleFunc("/{id:[0-9]+}", controllers.GetPositionHandler).Methods("GET")
	position.HandleFunc("", controllers.PostPositionHandler).Methods("POST")
	position.HandleFunc("/{id:[0-9]+}", controllers.PutPositionHandler).Methods("PUT")
	position.HandleFunc("/{id:[0-9]+}", controllers.DeletePositionHandler).Methods("DELETE")

	return R
}
