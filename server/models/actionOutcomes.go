package models

type ActionOutcome struct {
	ActionID  int64 `json:"action" binding:"required"`
	OutcomeID int64 `json:"outcome" binding:"required"`
}

func GetActionOutcomes(actionId int64) ([]int64, error) {
	rows, err := Database.Query("SELECT belief_id FROM action_outcomes WHERE action_id = ?", actionId)

	if err != nil {
		return nil, err
	}

	var outcomes []int64
	for rows.Next() {
		var outcome int64
		rows.Scan(&outcome)
		outcomes = append(outcomes, outcome)
	}

	return outcomes, nil
}

func AddActionOutcomes(actionOutcomes []ActionOutcome) error {
	for _, actionOutcome := range actionOutcomes {
		_, err := Database.Exec("INSERT INTO action_outcomes (action_id, belief_id) VALUES (?, ?)", actionOutcome.ActionID, actionOutcome.OutcomeID)

		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveActionOutcomes(actionId int64) error {
	_, err := Database.Exec("DELETE FROM action_outcomes WHERE action_id = ?", actionId)

	if err != nil {
		return err
	}

	return nil
}
