package controllers

import (
	"encoding/json"
	"net/http"

	m "github.com/alvesRenan/rest-golang-test/model"
	"github.com/alvesRenan/rest-golang-test/utils"
	"github.com/gorilla/mux"
)

var s m.Scenario

// CreateScenario adds new scenario
func CreateScenario(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload!")
		return
	}

	if err := s.CreateScenario(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, s)
}

// GetScenarios returns all scenarios
func GetScenarios(w http.ResponseWriter, r *http.Request) {
	scenarios, err := m.GetScenarios()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, scenarios)
}

// DeleteScenario deletes a scenario
func DeleteScenario(w http.ResponseWriter, r *http.Request) {
	s.Name = mux.Vars(r)["name"]
	if err := s.DeleteScenario(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
