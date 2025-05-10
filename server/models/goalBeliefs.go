package models

type GoalBelief struct {
	GoalID   int64 `json:"goal" binding:"required"`
	BeliefID int64 `json:"belief" binding:"required"`
}

func GetGoalBeliefs(goal_id int64) ([]int64, error) {
	rows, err := Database.Query("SELECT belief_id FROM goal_beliefs WHERE goal_id = ?", goal_id)

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

func AddGoalBeliefs(goalBeliefs []GoalBelief) error {
	for _, goalBelief := range goalBeliefs {
		_, err := Database.Exec("INSERT INTO goal_beliefs (goal_id, belief_id) VALUES (?, ?)", goalBelief.GoalID, goalBelief.BeliefID)

		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveGoalBeliefs(goal_id int64) error {
	_, err := Database.Exec("DELETE FROM goal_beliefs WHERE goal_id = ?", goal_id)

	if err != nil {
		return err
	}

	return nil
}
