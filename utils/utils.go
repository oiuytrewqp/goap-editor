package utils

var currentId int = 0

func GetId() int {
	currentId += 1

	return currentId
}
