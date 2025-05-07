package models

import "database/sql"

type Goal struct {
	ID   int64
	Name string `binding:"required"`
}

func GetGoals() ([]Goal, error) {
	sql.Open("sqlite3", "./database")
	return []Goal{}, nil
}

func GetGoal(id int) (Goal, error) {
	return Goal{}, nil
}

func CreateGoal(goal Goal) error {
	return nil
}

func UpdateGoal(id int, goal Goal) error {
	return nil
}

func DeleteGoal(id int) error {
	return nil
}
