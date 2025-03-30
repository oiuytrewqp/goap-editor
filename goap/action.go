package goap

import "oiuytrewqp/goap-editor/utils"

type Action struct {
	Id            int
	Name          string
	Prerequisites Beliefs
	Outcomes      Beliefs
	Action        string
	Location      string
}

func NewAction(name string) *Action {
	return &Action{
		Id:            utils.GetId(),
		Name:          name,
		Prerequisites: make(Beliefs),
		Outcomes:      make(Beliefs),
		Action:        "",
		Location:      "",
	}
}
