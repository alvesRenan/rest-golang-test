package controllers

import (
	"encoding/json"
	"net/http"

	m "github.com/alvesRenan/rest-golang-test/model"
	"github.com/alvesRenan/rest-golang-test/utils"
	"github.com/gorilla/mux"
)

var c m.Container

// CreateContainer adds new container
func CreateContainer(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload!")
		return
	}

	if err := c.CreateContainer(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, c)
}

// GetContainers returns all containers
func GetContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := m.GetContainers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, containers)
}

// DeleteContainer deletes a container
func DeleteContainer(w http.ResponseWriter, r *http.Request) {
	c.Name = mux.Vars(r)["name"]
	if err := c.DeleteContainer(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
