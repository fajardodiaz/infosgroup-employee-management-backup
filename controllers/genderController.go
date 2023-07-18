package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fajardodiaz/infosgroup-employee-management/initializer"
	"github.com/fajardodiaz/infosgroup-employee-management/models"
	"github.com/gorilla/mux"
)

func GetGendersHandler(w http.ResponseWriter, r *http.Request) {
	var genders []models.Gender
	initializer.Db.Find(&genders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&genders)
}

func GetGenderHandler(w http.ResponseWriter, r *http.Request) {
	var gender models.Gender
	vars := mux.Vars(r)

	initializer.Db.Find(&gender, vars["id"])
	if gender.ID == 0 {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Error, gender not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&gender)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&gender)
}

func PutGenderHandler(w http.ResponseWriter, r *http.Request) {
	var gender models.Gender
	var newGender models.Gender
	vars := mux.Vars(r)
	initializer.Db.Find(&gender, vars["id"])

	if gender.ID == 0 {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Error, gender not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewDecoder(r.Body).Decode(&newGender)
	err := newGender.Validate()

	if err != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	gender.Name = newGender.Name
	initializer.Db.Save(&gender)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&gender)
}

func DeleteGenderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var gender models.Gender
	err := initializer.Db.Where("id = ?", vars["id"]).Take(&gender).Delete(&models.Gender{})
	if err.Error != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gender)

}
