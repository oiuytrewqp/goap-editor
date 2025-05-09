package models

type Action struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Function string `json:"function"`
	Location int64  `json:"location"`
}

func GetActions() ([]Action, error) {
	rows, err := Database.Query("SELECT * FROM actions")

	if err != nil {
		return nil, err
	}

	var actions []Action
	for rows.Next() {
		var action Action
		rows.Scan(&action.ID, &action.Name, &action.Function, &action.Location)
		actions = append(actions, action)
	}

	return actions, nil
}

func GetAction(id int) (Action, error) {
	rows, err := Database.Query("SELECT * FROM actions WHERE id = ?", id)

	if err != nil {
		return Action{}, err
	}

	var action Action
	err = rows.Scan(&action.ID, &action.Name, &action.Function, &action.Location)

	if err != nil {
		return Action{}, err
	}

	return action, nil
}

func CreateAction(action Action) (int64, error) {
	result, err := Database.Exec("INSERT INTO actions (name, function, location) VALUES (?, ?, ?)", action.Name, action.Function, action.Location)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateAction(id int, action Action) error {
	_, err := Database.Exec("UPDATE actions SET name = ?, function = ?, location = ? WHERE id = ?", action.Name, action.Function, action.Location, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAction(id int) error {
	_, err := Database.Exec("DELETE FROM actions WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
