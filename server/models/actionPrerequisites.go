package models

type ActionPrerequisite struct {
	ActionID       int64 `json:"action" binding:"required"`
	PrerequisiteID int64 `json:"prerequisite" binding:"required"`
}

func GetActionPrerequisites(action_id int64) ([]int64, error) {
	rows, err := Database.Query("SELECT action_id FROM action_prerequisites WHERE action_id = ?", action_id)

	if err != nil {
		return nil, err
	}

	var prerequisites []int64
	for rows.Next() {
		var prerequisite int64
		rows.Scan(&prerequisite)
		prerequisites = append(prerequisites, prerequisite)
	}

	return prerequisites, nil
}

func AddActionPrerequisites(actionPrerequisites []ActionPrerequisite) error {
	for _, actionPrerequisite := range actionPrerequisites {
		_, err := Database.Exec("INSERT INTO action_prerequisites (action_id, belief_id) VALUES (?, ?)", actionPrerequisite.ActionID, actionPrerequisite.PrerequisiteID)

		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveActionPrerequisites(action_id int64) error {
	_, err := Database.Exec("DELETE FROM action_prerequisites WHERE action_id = ?", action_id)

	if err != nil {
		return err
	}

	return nil
}
