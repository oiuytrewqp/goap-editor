package models

type Action struct {
	ID       int64
	Name     string `binding:"required"`
	Function string
	Location int64
}

func GetActions() ([]Action, error) {
	rows, err := database.Query("SELECT * FROM actions")

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
	rows, err := database.Query("SELECT * FROM actions WHERE id = ?", id)

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
	result, err := database.Exec("INSERT INTO actions (name, function, location) VALUES (?, ?, ?)", action.Name, action.Function, action.Location)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateAction(id int, action Action) error {
	return nil
}

func DeleteAction(id int) error {
	return nil
}
