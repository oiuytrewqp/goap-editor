package goap

import "oiuytrewqp/goap-editor/utils"

type Goal struct {
	Id       int
	Name     string
	Outcomes map[int]int
}

func NewGoal(name string) *Goal {
	return &Goal{
		Id:   utils.GetId(),
		Name: name,
	}
}
