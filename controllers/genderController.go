package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fajardodiaz/infosgroup-employee-management/initializer"
	"github.com/fajardodiaz/infosgroup-employee-management/models"
	"github.com/gorilla/mux"
)

func GetGendersHandler(w http.ResponseWriter, r *http.Request) {
	var genders []models.Gender
	initializer.Db.Find(&genders)
	response, err := json.Marshal(&genders)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(response))
}

func GetGenderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var gender models.Gender
	initializer.Db.First(&gender, id)

}

func PostGenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST"))
}

func PutGenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PUT"))
}

func DeleteGenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE"))
}
