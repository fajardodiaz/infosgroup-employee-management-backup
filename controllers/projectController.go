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

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	// Controller to get all projects
	var project []models.Project
	initializer.Db.Find(&project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&project)
}

func GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	// Controller to get a single project
	vars := mux.Vars(r)
	var project models.Project

	result := initializer.Db.Where(&project.ID, vars["id"]).First(&project)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		w.WriteHeader(http.StatusNotFound)
		resp["message"] = "Error, project not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Return the project
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&project)
}

func PostProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	json.NewDecoder(r.Body).Decode(&project)

	// This function validate if the project is valide
	err := project.Validate()
	if err != nil {
		resp := make(map[string]string)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	// This code create the project after the validation
	initializer.Db.Create(&project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&project)
}

func PutProjectHandler(w http.ResponseWriter, r *http.Request) {
	// this variable is the found project in the db
	var project models.Project

	// This variable is the new item
	var newProject models.Project

	vars := mux.Vars(r)

	result := initializer.Db.Where(&project.ID, vars["id"]).First(&project)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Error, project not found"
		json.NewEncoder(w).Encode(resp)
		return
	}

	// encode the new item
	json.NewDecoder(r.Body).Decode(&newProject)
	err := newProject.Validate()

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
	project.Name = newProject.Name
	initializer.Db.Save(&project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&project)
}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var project models.Project
	// Find if the item exists
	err := initializer.Db.Where("id = ?", vars["id"]).Take(&project).Delete(&models.Project{})
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
	json.NewEncoder(w).Encode(project)
}
