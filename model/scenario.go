package model

import "github.com/alvesRenan/rest-golang-test/conf"

// Scenario is a struct that has all components of a scenario
type Scenario struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

// CreateScenario add a new scenario to the db
func (s *Scenario) CreateScenario() error {
	db := conf.ConnectDB()

	stmt, _ := db.Prepare("INSERT INTO scenarios (name, state) VALUES (?,?)")
	defer stmt.Close()

	if _, err := stmt.Exec(s.Name, s.State); err != nil {
		return nil
	}

	defer db.Close()
	return nil
}

// DeleteScenario deletes a scenario given a name
func (s *Scenario) DeleteScenario() error {
	db := conf.ConnectDB()
	_, err := db.Exec("DELETE FROM scenarios WHERE name=$1", s.Name)

	defer db.Close()
	return err
}

// GetScenarios returns a list of all scenarios
func GetScenarios() ([]Scenario, error) {
	db := conf.ConnectDB()
	rows, err := db.Query("SELECT * FROM scenarios")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	scenarios := []Scenario{}
	var s Scenario

	for rows.Next() {
		if err := rows.Scan(&s.Name, &s.State); err != nil {
			return nil, err
		}

		scenarios = append(scenarios, s)
	}

	defer db.Close()
	return scenarios, nil
}
