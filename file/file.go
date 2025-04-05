package file

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

//try to open file and if failed then return object
//save object and if failed then try to create directory

const mkDirPermission = 0755

func Load[T any](path string, filename string) (T, error) {
	fullPath := filepath.Join(path, filename)

	jsonFile, err := os.Open(fullPath)

	var object T

	if err != nil {
		return object, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return object, err
	}

	json.Unmarshal(byteValue, &object)

	return object, nil
}

func Save[T any](path string, filename string, object T) error {
	fullPath := filepath.Join(path, filename)

	os.Mkdir(path, mkDirPermission)

	byteValue, err := json.MarshalIndent(object, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, byteValue, mkDirPermission)

	return err
}
