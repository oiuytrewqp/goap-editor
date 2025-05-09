package models

type Belief struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" binding:"required"`
	Value int64  `json:"value"`
}

func GetBeliefs() ([]Belief, error) {
	rows, err := Database.Query("SELECT * FROM beliefs")

	if err != nil {
		return nil, err
	}

	var beliefs []Belief
	for rows.Next() {
		var belief Belief
		rows.Scan(&belief.ID, &belief.Name, &belief.Value)
		beliefs = append(beliefs, belief)
	}

	return beliefs, nil
}

func GetBelief(id int) (Belief, error) {
	rows, err := Database.Query("SELECT * FROM beliefs WHERE id = ?", id)

	if err != nil {
		return Belief{}, err
	}

	var belief Belief
	err = rows.Scan(&belief.ID, &belief.Name, &belief.Value)

	if err != nil {
		return Belief{}, err
	}

	return belief, nil
}

func CreateBelief(belief Belief) (int64, error) {
	result, err := Database.Exec("INSERT INTO beliefs (name, amount) VALUES (?, ?)", belief.Name, belief.Value)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateBelief(id int, belief Belief) error {
	_, err := Database.Exec("UPDATE beliefs SET name = ?, amount = ? WHERE id = ?", belief.Name, belief.Value, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBelief(id int) error {
	_, err := Database.Exec("DELETE FROM beliefs WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
