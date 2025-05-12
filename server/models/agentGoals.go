package models

type AgentGoal struct {
	AgentID int64 `json:"agent" binding:"required"`
	GoalID  int64 `json:"goal" binding:"required"`
}

func GetAgentGoals(agentId int64) ([]int64, error) {
	rows, err := Database.Query("SELECT goal_id FROM agent_goals WHERE agent_id = ?", agentId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var goals []int64
	for rows.Next() {
		var goal int64
		rows.Scan(&goal)
		goals = append(goals, goal)
	}

	return goals, nil
}

func AddAgentGoals(agentGoals []AgentGoal) error {
	for _, agentGoal := range agentGoals {
		_, err := Database.Exec("INSERT INTO agent_goals (agent_id, goal_id) VALUES (?, ?)", agentGoal.AgentID, agentGoal.GoalID)

		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveAgentGoals(agentId int64) error {
	_, err := Database.Exec("DELETE FROM agent_goals WHERE agent_id = ?", agentId)

	if err != nil {
		return err
	}

	return nil
}
