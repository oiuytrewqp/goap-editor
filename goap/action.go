package goap

import "oiuytrewqp/goap-editor/utils"

type Action struct {
	Id            int
	Name          string
	Prerequisites map[int]int
	Outcomes      map[int]int
	Action        int
	Location      int
}

func NewAction(name string) *Action {
	return &Action{
		Id:   utils.GetId(),
		Name: name,
	}
}
