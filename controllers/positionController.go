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

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	// Controller to get all positions
	var position []models.Position
	initializer.Db.Find(&position)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&position)
}

func GetPositionHandler(w http.ResponseWriter, r *http.Request) {
	// Controller to get a single position
	vars := mux.Vars(r)
	var position models.Position
	initializer.Db.First(&position, vars["id"])
	// Validate if the position ID == 0
	if position.ID == 0 {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Error, position not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Return the position
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&position)
}

func PostPositionHandler(w http.ResponseWriter, r *http.Request) {
	var position models.Position
	json.NewDecoder(r.Body).Decode(&position)

	// This function validate if the position is valide
	err := position.Validate()
	if err != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	// This code create the position after the validation
	initializer.Db.Create(&position)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&position)
}

func PutPositionHandler(w http.ResponseWriter, r *http.Request) {
	// this variable is the found position in the db
	var position models.Position

	// This variable is the new item
	var newPosition models.Position

	vars := mux.Vars(r)

	result := initializer.Db.Where(&position.ID, vars["id"]).First(&position)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Error, position not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// encode the new item
	json.NewDecoder(r.Body).Decode(&newPosition)
	err := newPosition.Validate()

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
	position.Name = newPosition.Name
	initializer.Db.Save(&position)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&position)
}

func DeletePositionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var position models.Position
	// Find if the item exists
	err := initializer.Db.Where("id = ?", vars["id"]).Take(&position).Delete(&models.Position{})
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(position)
}
