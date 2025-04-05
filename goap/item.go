package goap

import "oiuytrewqp/goap-editor/utils"

type Item struct {
	Id   int
	name string
}

func NewItem(name string) *Item {
	return &Item{
		Id:   utils.GetId(),
		name: name,
	}
}
