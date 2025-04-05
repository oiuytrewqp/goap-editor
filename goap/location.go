package goap

import "oiuytrewqp/goap-editor/utils"

type Location struct {
	Id   int
	Name string
}

func NewLocation(name string) *Location {
	return &Location{
		Id:   utils.GetId(),
		Name: name,
	}
}
