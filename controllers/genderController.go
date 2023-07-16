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
	var user models.Gender
	vars := mux.Vars(r)
	initializer.Db.Find(&user, vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&user)
}

func PostGenderHandler(w http.ResponseWriter, r *http.Request) {
	var gender models.Gender
	json.NewDecoder(r.Body).Decode(&gender)

	err := gender.Validate()
	if err != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	initializer.Db.Create(&gender)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&gender)
}

func PutGenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PUT"))
}

func DeleteGenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE"))
}
