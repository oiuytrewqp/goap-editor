package models

type Goal struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
}

func GetGoals() ([]Goal, error) {
	rows, err := Database.Query("SELECT * FROM goals")

	if err != nil {
		return nil, err
	}

	var goals []Goal
	for rows.Next() {
		var goal Goal
		rows.Scan(&goal.ID, &goal.Name)
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
	err = rows.Scan(&goal.ID, &goal.Name)

	if err != nil {
		return Goal{}, err
	}

	return goal, nil
}

func CreateGoal(goal Goal) (int64, error) {
	result, err := Database.Exec("INSERT INTO goals (name) VALUES (?)", goal.Name)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateGoal(id int, goal Goal) error {
	_, err := Database.Exec("UPDATE goals SET name = ? WHERE id = ?", goal.Name, id)

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
