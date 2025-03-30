package goap

type Goap struct {
	World     Agent
	Locations []string
	items     []string
	actions   map[string]Action
	Agents    []Agent
}

func NewGoap() Goap {
	newGoap := Goap{}

	newGoap.World = *NewAgent("World")

	newGoap.Locations = make([]string, 5)
	newGoap.Locations[0] = "Home"
	newGoap.Locations[1] = "Enterance"
	newGoap.Locations[2] = "Reception"
	newGoap.Locations[3] = "Waiting Room"
	newGoap.Locations[4] = "Break Room"

	newGoap.items = make([]string, 3)
	newGoap.items[0] = "Booth 1"
	newGoap.items[1] = "Booth 2"
	newGoap.items[2] = "Booth 3"

	newGoap.actions = make(map[string]Action)

	var goToHospital = *NewAction("Go to Hospital")
	goToHospital.Outcomes["atEnterance"] = 1

	newGoap.actions["Go to Hospital"] = goToHospital

	var registerAtReception = *NewAction("Register at Reception")
	registerAtReception.Prerequisites["atEnterance"] = 1
	registerAtReception.Outcomes["resistered"] = 1

	newGoap.actions["Register at Reception"] = registerAtReception

	var waitInWaitingRoom = *NewAction("Wait in Waiting Room")
	waitInWaitingRoom.Prerequisites["resistered"] = 1
	waitInWaitingRoom.Outcomes["waitingInWaitingRoom"] = 1

	newGoap.actions["Wait in Waiting Room"] = waitInWaitingRoom

	var goHome = *NewAction("Go Home")
	goHome.Prerequisites["waiting"] = 1

	newGoap.actions["Go Home"] = goHome

	var waitInBreakRoom = *NewAction("Wait in Break Room")
	waitInBreakRoom.Prerequisites["atEnterance"] = 1
	waitInBreakRoom.Outcomes["waitingInBreakRoom"] = 1

	var nurseActions = make([]Action, 2)
	nurseActions[0] = goToHospital
	nurseActions[1] = waitInBreakRoom

	var pateintActions = make([]Action, 4)
	pateintActions[0] = goToHospital
	pateintActions[1] = registerAtReception
	pateintActions[2] = waitInWaitingRoom
	pateintActions[3] = goHome

	newGoap.Agents = make([]Agent, 8)
	newGoap.Agents[0] = *newGoap.newNurse("Nurse 1")
	newGoap.Agents[1] = *newGoap.newNurse("Nurse 2")
	newGoap.Agents[2] = *newGoap.newPatient("Pateint 1")
	newGoap.Agents[3] = *newGoap.newPatient("Pateint 2")
	newGoap.Agents[4] = *newGoap.newPatient("Pateint 5")
	newGoap.Agents[5] = *newGoap.newPatient("Pateint 6")
	newGoap.Agents[6] = *newGoap.newPatient("Pateint 7")
	newGoap.Agents[7] = *newGoap.newPatient("Pateint 8")

	return newGoap
}

func (goap *Goap) newNurse(name string) *Agent {
	var nurse = NewAgent(name)
	nurse.Actions = []Action{
		goap.actions["Go to Hospital"],
		goap.actions["Wait in Break Room"],
	}

	return nurse
}

func (goap *Goap) newPatient(name string) *Agent {
	var patient = NewAgent(name)
	patient.Actions = []Action{
		goap.actions["Go to Hospital"],
		goap.actions["Register at Reception"],
		goap.actions["Wait in Waiting Room"],
		goap.actions["Go Home"],
	}

	return patient
}
