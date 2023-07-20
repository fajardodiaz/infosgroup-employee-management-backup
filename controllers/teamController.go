package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fajardodiaz/infosgroup-employee-management/initializer"
	"github.com/fajardodiaz/infosgroup-employee-management/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	// Controller to get all teams
	var team []models.Team
	initializer.Db.Find(&team)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&team)
}

func GetTeamHandler(w http.ResponseWriter, r *http.Request) {
	// Controller to get a single team
	vars := mux.Vars(r)
	var team models.Team

	result := initializer.Db.Where(&team.ID, vars["id"]).First(&team)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		w.WriteHeader(http.StatusNotFound)
		resp["message"] = "Error, team not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Return the team
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&team)
}

func PostTeamHandler(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	json.NewDecoder(r.Body).Decode(&team)

	// This function validate if the team is valide
	err := team.Validate()
	if err != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	// This code create the team after the validation
	initializer.Db.Create(&team)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&team)
}

func PutTeamHandler(w http.ResponseWriter, r *http.Request) {
	// this variable is the found team in the db
	var team models.Team

	// This variable is the new item
	var newTeam models.Team

	vars := mux.Vars(r)

	result := initializer.Db.Where(&team.ID, vars["id"]).First(&team)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Error, team not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// encode the new item
	json.NewDecoder(r.Body).Decode(&newTeam)
	err := newTeam.Validate()

	// // Validate if the item is correct
	if err != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	// // Save the new item
	team.Name = newTeam.Name
	initializer.Db.Save(&team)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&team)
}

func DeleteTeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var team models.Team
	// Find if the item exists
	err := initializer.Db.Where("id = ?", vars["id"]).Take(&team).Delete(&models.Team{})
	if err.Error != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	// return the deleted item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(team)
}
