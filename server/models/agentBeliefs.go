package models

type AgentBelief struct {
	AgentID  int64 `json:"agent" binding:"required"`
	BeliefID int64 `json:"belief" binding:"required"`
}

func GetAgentBeliefs(agentId int64) ([]int64, error) {
	rows, err := Database.Query("SELECT belief_id FROM agent_beliefs WHERE agent_id = ?", agentId)

	if err != nil {
		return nil, err
	}

	var beliefs []int64
	for rows.Next() {
		var belief int64
		rows.Scan(&belief)
		beliefs = append(beliefs, belief)
	}

	return beliefs, nil
}

func AddAgentBeliefs(agentBeliefs []AgentBelief) error {
	for _, agentBelief := range agentBeliefs {
		_, err := Database.Exec("INSERT INTO agent_beliefs (agent_id, belief_id) VALUES (?, ?)", agentBelief.AgentID, agentBelief.BeliefID)

		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveAgentBeliefs(agentId int64) error {
	_, err := Database.Exec("DELETE FROM agent_beliefs WHERE agent_id = ?", agentId)

	if err != nil {
		return err
	}

	return nil
}
