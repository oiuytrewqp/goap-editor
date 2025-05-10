package models

type Goal struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Priority    int64   `json:"priority" binding:"required"`
	Beliefs     []int64 `json:"beliefs" binding:"required"`
	Description string  `json:"description"`
}

func GetGoals() ([]Goal, error) {
	rows, err := Database.Query("SELECT * FROM goals")

	if err != nil {
		return nil, err
	}

	var goals []Goal
	for rows.Next() {
		var goal Goal
		rows.Scan(&goal.ID, &goal.Name, &goal.Priority, &goal.Description)
		goals = append(goals, goal)
	}

	return goals, nil
}

func GetGoal(id int) (Goal, error) {
	rows, err := Database.Query("SELECT * FROM goals WHERE id = ?", id)

	if err != nil {
		return Goal{}, err
	}

	var goal Goal
	err = rows.Scan(&goal.ID, &goal.Name, &goal.Priority, &goal.Description)

	if err != nil {
		return Goal{}, err
	}

	return goal, nil
}

func CreateGoal(goal Goal) (int64, error) {
	result, err := Database.Exec("INSERT INTO goals (name, priority, description) VALUES (?, ?, ?)", goal.Name, goal.Priority, goal.Description)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateGoal(id int, goal Goal) error {
	_, err := Database.Exec("UPDATE goals SET name = ?, priority = ?, description = ? WHERE id = ?", goal.Name, goal.Priority, goal.Description, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteGoal(id int) error {
	_, err := Database.Exec("DELETE FROM goals WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
