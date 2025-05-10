package models

type Action struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name" binding:"required"`
	Method        string  `json:"method"`
	LocationID    int64   `json:"locationId"`
	Prerequisites []int64 `json:"prerequisites"`
	Outcomes      []int64 `json:"outcomes"`
	Description   string  `json:"description"`
}

func GetActions() ([]Action, error) {
	rows, err := Database.Query("SELECT * FROM actions")

	if err != nil {
		return nil, err
	}

	var actions []Action
	for rows.Next() {
		var action Action
		rows.Scan(&action.ID, &action.Name, &action.Method, &action.LocationID, &action.Description)
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
	err = rows.Scan(&action.ID, &action.Name, &action.Method, &action.LocationID, &action.Description)

	if err != nil {
		return Action{}, err
	}

	return action, nil
}

func CreateAction(action Action) (int64, error) {
	result, err := Database.Exec("INSERT INTO actions (name, method, location, description) VALUES (?, ?, ?, ?)", action.Name, action.Method, action.LocationID, action.Description)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateAction(id int, action Action) error {
	_, err := Database.Exec("UPDATE actions SET name = ?, method = ?, location = ?, description = ? WHERE id = ?", action.Name, action.Method, action.LocationID, action.Description, id)

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
