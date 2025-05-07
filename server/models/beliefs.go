package models

import "database/sql"

type Belief struct {
	ID   int64
	Name string `binding:"required"`
}

func GetBeliefs() ([]Belief, error) {
	sql.Open("sqlite3", "./database")
	return []Belief{}, nil
}

func GetBelief(id int) (Belief, error) {
	return Belief{}, nil
}

func CreateBelief(belief Belief) error {
	return nil
}

func UpdateBelief(id int, belief Belief) error {
	return nil
}

func DeleteBelief(id int) error {
	return nil
}
