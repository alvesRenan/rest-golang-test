package main

import "database/sql"

// Scenario is a struct that has all components of a scenario
type Scenario struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

func (s *Scenario) createScenario(db *sql.DB) error {
	stmt, _ := db.Prepare("INSERT INTO scenarios (name, state) VALUES (?,?)")
	defer stmt.Close()

	if _, err := stmt.Exec(s.Name, s.State); err != nil {
		return nil
	}

	return nil
}

func (s *Scenario) deleteScenario(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM scenarios WHERE name=$1", s.Name)

	return err
}

func getScenarios(db *sql.DB) ([]Scenario, error) {
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

	return scenarios, nil
}
