package models

type Agent struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func GetAgents() ([]Agent, error) {
	rows, err := Database.Query("SELECT * FROM agents")

	if err != nil {
		return nil, err
	}

	var agents []Agent
	for rows.Next() {
		var agent Agent
		rows.Scan(&agent.ID, &agent.Name, &agent.Description)
		agents = append(agents, agent)
	}

	return agents, nil
}

func GetAgent(id int) (Agent, error) {
	rows, err := Database.Query("SELECT * FROM agents WHERE id = ?", id)

	if err != nil {
		return Agent{}, err
	}

	var agent Agent
	err = rows.Scan(&agent.ID, &agent.Name, &agent.Description)

	if err != nil {
		return Agent{}, err
	}

	return agent, nil
}

func CreateAgent(agent Agent) (int64, error) {
	result, err := Database.Exec("INSERT INTO agents (name, description) VALUES (?, ?)", agent.Name, agent.Description)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateAgent(id int, agent Agent) error {
	_, err := Database.Exec("UPDATE agents SET name = ?, description = ? WHERE id = ?", agent.Name, agent.Description, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAgent(id int) error {
	_, err := Database.Exec("DELETE FROM agents WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
