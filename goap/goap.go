package goap

import (
	"oiuytrewqp/goap-editor/file"
)

type Goap struct {
	World       int
	WorldAgents []int
	Agents      map[int]Agent
	Locations   map[int]Location
	items       map[int]Item
	Beliefs     map[int]Belief
	Actions     map[int]Action
	Goals       map[int]Goal
}

func NewGoap() Goap {
	newGoap := Goap{
		Agents:    make(map[int]Agent),
		Locations: make(map[int]Location),
		items:     make(map[int]Item),
		Beliefs:   make(map[int]Belief),
		Actions:   make(map[int]Action),
		Goals:     make(map[int]Goal),
	}

	atEnterance := newGoap.createBelief("atEnterance")
	resistered := newGoap.createBelief("resistered")
	waitingInWaitingRoom := newGoap.createBelief("waitingInWaitingRoom")
	waiting := newGoap.createBelief("waiting")
	waitingInBreakRoom := newGoap.createBelief("waitingInBreakRoom")

	goToHospital := newGoap.createAction("Go to Hospital", nil, map[int]int{atEnterance: 1})
	goResisteredAtReception := newGoap.createAction("Register at Reception", map[int]int{atEnterance: 1}, map[int]int{resistered: 1})
	goWaitingInWaitingRoom := newGoap.createAction("Wait in Waiting Room", map[int]int{resistered: 1}, map[int]int{waitingInWaitingRoom: 1})
	goHome := newGoap.createAction("Go Home", map[int]int{waiting: 1}, nil)
	goWaitingInBreakRoom := newGoap.createAction("Wait in Break Room", map[int]int{atEnterance: 1}, map[int]int{waitingInBreakRoom: 1})

	newGoap.createLocation("Home")
	newGoap.createLocation("Enterance")
	newGoap.createLocation("Reception")
	newGoap.createLocation("Waiting Room")
	newGoap.createLocation("Break Room")

	newGoap.createItem("Booth 1")
	newGoap.createItem("Booth 2")
	newGoap.createItem("Booth 3")

	newGoap.World = newGoap.addAgent("Hospital", nil)

	newGoap.addAgent("Nurse 1", []int{goToHospital, goWaitingInBreakRoom})
	newGoap.addAgent("Nurse 2", []int{goToHospital, goWaitingInBreakRoom})
	newGoap.addAgent("Nurse 3", []int{goToHospital, goWaitingInBreakRoom})

	newGoap.addAgent("Patient 1", []int{goToHospital, goResisteredAtReception, goWaitingInWaitingRoom, goHome})
	newGoap.addAgent("Patient 2", []int{goToHospital, goResisteredAtReception, goWaitingInWaitingRoom, goHome})
	newGoap.addAgent("Patient 3", []int{goToHospital, goResisteredAtReception, goWaitingInWaitingRoom, goHome})
	newGoap.addAgent("Patient 4", []int{goToHospital, goResisteredAtReception, goWaitingInWaitingRoom, goHome})
	newGoap.addAgent("Patient 5", []int{goToHospital, goResisteredAtReception, goWaitingInWaitingRoom, goHome})
	newGoap.addAgent("Patient 6", []int{goToHospital, goResisteredAtReception, goWaitingInWaitingRoom, goHome})

	newGoap.WorldAgents = make([]int, len(newGoap.Agents)-1)
	i := 0
	for _, agent := range newGoap.Agents {
		if agent.Id != newGoap.World {
			newGoap.WorldAgents[i] = agent.Id
			i++
		}
	}

	file.Save("data", "test.json", newGoap)

	return newGoap
}

func (goap Goap) addAgent(name string, actions []int) int {
	newAgent := NewAgent(name)
	newAgent.Actions = actions
	goap.Agents[newAgent.Id] = *newAgent

	return newAgent.Id
}

func (goap Goap) createBelief(name string) int {
	newBelief := *NewBelief(name)
	goap.Beliefs[newBelief.Id] = newBelief

	return newBelief.Id
}

func (goap Goap) createAction(name string, prerequisites map[int]int, outcomes map[int]int) int {
	newAction := *NewAction(name)
	newAction.Prerequisites = prerequisites
	newAction.Outcomes = outcomes
	goap.Actions[newAction.Id] = newAction

	return newAction.Id
}

func (goap Goap) createLocation(name string) {
	newLocation := *NewLocation(name)
	goap.Locations[newLocation.Id] = newLocation
}

func (goap Goap) createItem(name string) {
	newItem := *NewItem(name)
	goap.items[newItem.Id] = newItem
}
