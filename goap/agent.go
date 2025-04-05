package goap

import (
	"encoding/json"
	"oiuytrewqp/goap-editor/utils"
)

type Agent struct {
	Id      int
	Name    string
	Beliefs map[int]int
	Actions []int
	Goals   []int
}

func NewAgent(name string) *Agent {
	return &Agent{
		Id:   utils.GetId(),
		Name: name,
	}
}

func (agent Agent) ToJson() (string, error) {
	json, err := json.Marshal(agent)

	if err != nil {
		return "", err
	}

	return string(json), nil
}
