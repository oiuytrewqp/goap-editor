package models

import "database/sql"

type Location struct {
	ID   int64
	Name string `binding:"required"`
}

func GetLocations() ([]Location, error) {
	sql.Open("sqlite3", "./database")
	return []Location{}, nil
}

func GetLocation(id int) (Location, error) {
	return Location{}, nil
}

func CreateLocation(location Location) error {
	return nil
}

func UpdateLocation(id int, location Location) error {
	return nil
}

func DeleteLocation(id int) error {
	return nil
}
