package model

import "github.com/alvesRenan/rest-golang-test/conf"

// Container is a struct that has all components of a container
type Container struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Network    string `json:"net"`
	AdbPort    string `json:"adb_port"`
	SerialPort string `json:"serial_port"`
	IsServer   int    `json:"is_server"`
	VNCPort    string `json:"vnc_port"`
	State      string `json:"state"`
}

// CreateContainer add a new scenario to the db
func (c *Container) CreateContainer() error {
	db := conf.ConnectDB()

	stmt, _ := db.Prepare("INSERT INTO containers (name, network, adb_port, serial_port, vnc_port) VALUES (?,?,?,?,?)")
	_, err := stmt.Exec(c.Name, c.Network, c.AdbPort, c.SerialPort, c.VNCPort)

	if err != nil {
		return nil
	}

	defer stmt.Close()
	defer db.Close()
	return nil
}

// DeleteContainer deletes a scenario given a name
func (c *Container) DeleteContainer() error {
	db := conf.ConnectDB()
	_, err := db.Exec("DELETE FROM containers WHERE name=$1", c.Name)

	defer db.Close()
	return err
}

// GetContainers returns a list of all containers
func GetContainers() ([]Container, error) {
	db := conf.ConnectDB()
	rows, err := db.Query("SELECT * FROM containers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	contaienrs := []Container{}
	var c Container

	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Name, &c.Network, &c.AdbPort, &c.SerialPort, &c.VNCPort, &c.IsServer, &c.State)
		if err != nil {
			return nil, err
		}

		contaienrs = append(contaienrs, c)
	}

	defer db.Close()
	return contaienrs, nil
}
