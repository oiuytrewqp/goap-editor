package models

type Location struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
}

func GetLocations() ([]Location, error) {
	rows, err := Database.Query("SELECT * FROM locations")

	if err != nil {
		return nil, err
	}

	var locations []Location
	for rows.Next() {
		var location Location
		rows.Scan(&location.ID, &location.Name)
		locations = append(locations, location)
	}

	return locations, nil
}

func GetLocation(id int) (Location, error) {
	rows, err := Database.Query("SELECT * FROM locations WHERE id = ?", id)

	if err != nil {
		return Location{}, err
	}

	var location Location
	err = rows.Scan(&location.ID, &location.Name)

	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func CreateLocation(location Location) (int64, error) {
	result, err := Database.Exec("INSERT INTO locations (name) VALUES (?)", location.Name)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateLocation(id int, location Location) error {
	_, err := Database.Exec("UPDATE locations SET name = ? WHERE id = ?", location.Name, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteLocation(id int) error {
	_, err := Database.Exec("DELETE FROM locations WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
