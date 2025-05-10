package models

type Location struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func GetLocations() ([]Location, error) {
	rows, err := Database.Query("SELECT * FROM locations")

	if err != nil {
		return nil, err
	}

	var locations []Location
	for rows.Next() {
		var location Location
		rows.Scan(&location.ID, &location.Name, &location.Description)
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
	err = rows.Scan(&location.ID, &location.Name, &location.Description)

	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func CreateLocation(location Location) (int64, error) {
	result, err := Database.Exec("INSERT INTO locations (name, description) VALUES (?, ?)", location.Name, location.Description)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func UpdateLocation(id int, location Location) error {
	_, err := Database.Exec("UPDATE locations SET name = ?, description = ? WHERE id = ?", location.Name, location.Description, id)

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
