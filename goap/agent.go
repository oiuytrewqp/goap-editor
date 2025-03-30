package goap

import (
	"encoding/json"
	"oiuytrewqp/goap-editor/utils"
)

type Agent struct {
	Id      int
	Name    string
	Beliefs map[string]int
	Actions []Action
	Goals   []Goal
}

type agentJson struct {
	Id      int
	Name    string
	Beliefs map[string]int
	Actions []int
	Goals   []int
}

func NewAgent(name string) *Agent {
	return &Agent{
		Id:      utils.GetId(),
		Name:    name,
		Beliefs: make(map[string]int),
		Actions: []Action{},
		Goals:   []Goal{},
	}
}

func (agent Agent) ToJson() (string, error) {
	var agentJson agentJson
	agentJson.Id = agent.Id
	agentJson.Name = agent.Name
	agentJson.Beliefs = agent.Beliefs
	agentJson.Actions = make([]int, len(agent.Actions))
	for i, action := range agent.Actions {
		agentJson.Actions[i] = action.Id
	}
	agentJson.Goals = make([]int, len(agent.Goals))
	for i, goal := range agent.Goals {
		agentJson.Goals[i] = goal.Id
	}

	json, err := json.Marshal(agentJson)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
