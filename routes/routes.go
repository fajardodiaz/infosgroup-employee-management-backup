package routes

import (
	"github.com/fajardodiaz/infosgroup-employee-management/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	R := mux.NewRouter()

	v1 := R.PathPrefix("/api/v1").Subrouter()

	// Gender routes
	genders := v1.PathPrefix("/genders").Subrouter()

	genders.HandleFunc("", controllers.GetGendersHandler).Methods("GET")
	genders.HandleFunc("/{id:[0-9]+}", controllers.GetGenderHandler).Methods("GET")
	genders.HandleFunc("", controllers.PostGenderHandler).Methods("POST")
	genders.HandleFunc("/{id}", controllers.PutGenderHandler).Methods("PUT")
	genders.HandleFunc("/{id}", controllers.DeleteGenderHandler).Methods("DELETE")

	// Position routes
	positions := v1.PathPrefix("/positions").Subrouter()
	positions.HandleFunc("", controllers.GetPositionsHandler).Methods("GET")
	positions.HandleFunc("/{id:[0-9]+}", controllers.GetPositionHandler).Methods("GET")
	positions.HandleFunc("", controllers.PostPositionHandler).Methods("POST")
	positions.HandleFunc("/{id:[0-9]+}", controllers.PutPositionHandler).Methods("PUT")
	positions.HandleFunc("/{id:[0-9]+}", controllers.DeletePositionHandler).Methods("DELETE")

	// Project Routes
	projects := v1.PathPrefix("/projects").Subrouter()
	projects.HandleFunc("", controllers.GetProjectsHandler).Methods("GET")
	projects.HandleFunc("/{id:[0-9]+}", controllers.GetProjectHandler).Methods("GET")
	projects.HandleFunc("", controllers.PostProjectHandler).Methods("POST")
	projects.HandleFunc("/{id:[0-9]+}", controllers.PutProjectHandler).Methods("PUT")
	projects.HandleFunc("/{id:[0-9]+}", controllers.DeleteProjectHandler).Methods("DELETE")

	// State Routes
	states := v1.PathPrefix("/states").Subrouter()

	// Team Position
	teams := v1.PathPrefix("/teams").Subrouter()
	teams.HandleFunc("", controllers.GetTeamsHandler).Methods("GET")
	teams.HandleFunc("/{id:[0-9]+}", controllers.GetTeamHandler).Methods("GET")
	teams.HandleFunc("", controllers.PostTeamHandler).Methods("POST")
	teams.HandleFunc("/{id:[0-9]+}", controllers.PutTeamHandler).Methods("PUT")
	teams.HandleFunc("/{id:[0-9]+}", controllers.DeleteTeamHandler).Methods("DELETE")

	// Employee Position
	emplooyes := v1.PathPrefix("/employees").Subrouter()

	return R
}
