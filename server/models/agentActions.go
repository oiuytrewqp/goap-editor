package models

type AgentAction struct {
	AgentID  int64 `json:"agent" binding:"required"`
	ActionID int64 `json:"action" binding:"required"`
}

func GetAgentActions(agentId int64) ([]int64, error) {
	rows, err := Database.Query("SELECT action_id FROM agent_actions WHERE agent_id = ?", agentId)

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

func AddAgentActions(agentActions []AgentAction) error {
	for _, agentAction := range agentActions {
		_, err := Database.Exec("INSERT INTO agent_actions (agent_id, action_id) VALUES (?, ?)", agentAction.AgentID, agentAction.ActionID)

		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveAgentActions(agentId int64) error {
	_, err := Database.Exec("DELETE FROM agent_actions WHERE agent_id = ?", agentId)

	if err != nil {
		return err
	}

	return nil
}
