package main

import (
	"log"
	"net/http"

	"github.com/fajardodiaz/infosgroup-employee-management/initializer"
	"github.com/fajardodiaz/infosgroup-employee-management/routes"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
}

func main() {
	r := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":8000", r))
}
