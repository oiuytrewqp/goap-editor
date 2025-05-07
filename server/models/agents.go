package models

import "database/sql"

type Agent struct {
	ID   int64
	Name string `binding:"required"`
}

func GetAgents() ([]Agent, error) {
	sql.Open("sqlite3", "./database")
	return []Agent{}, nil
}

func GetAgent(id int) (Agent, error) {
	return Agent{}, nil
}

func CreateAgent(agent Agent) error {
	return nil
}

func UpdateAgent(id int, agent Agent) error {
	return nil
}

func DeleteAgent(id int) error {
	return nil
}
