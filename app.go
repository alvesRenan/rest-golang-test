package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// App is the main application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

var s Scenario

// Initialize create the DB connection and register the routes
func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", "./example-db.db")
	if err != nil {
		panic(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()

	a.initializeDB()
}

// routes
func (a *App) initializeRoutes() {
	// scenario routes
	a.Router.HandleFunc("/scneario/list", a.getScenarios).Methods("GET")
	a.Router.HandleFunc("/scenario/create", a.createScenario).Methods("POST")
	a.Router.HandleFunc("/scenario/delete/{name}", a.deleteScenario).Methods("DELETE")

	// container routes
	a.Router.HandleFunc("/container/list", getDevices).Methods("GET")
	a.Router.HandleFunc("/container/create", createContainer).Methods("POST")
	a.Router.HandleFunc("/container/delete/{id}", deleteDevice).Methods("DELETE")
	a.Router.HandleFunc("/container/update/{id}", updateDevice).Methods("UPDATE")
}

// initialize the tables
func (a *App) initializeDB() {
	sql := `CREATE TABLE IF NOT EXISTS scenarios (
		name TEXT PRIMARY KEY,
		state TEXT)`
	sts, _ := a.DB.Prepare(sql)
	sts.Exec()
}

// Run starts the server
func (a *App) Run() {
	log.Println("Starting server at localhost")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

// configures the responses
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// sends error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) createScenario(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload!")
		return
	}

	if err := s.createScenario(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, s)
}

func (a *App) getScenarios(w http.ResponseWriter, r *http.Request) {
	scenarios, err := getScenarios(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, scenarios)
}

func (a *App) deleteScenario(w http.ResponseWriter, r *http.Request) {
	s.Name = mux.Vars(r)["name"]
	if err := s.deleteScenario(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
