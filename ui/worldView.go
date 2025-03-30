package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type WorldView struct {
	theme         material.Theme
	world         goap.Agent
	worldButton   AgentButton
	agents        []goap.Agent
	agentsButtons []AgentButton
	Selected      goap.Agent
}

func NewWorldView(theme *material.Theme, world *goap.Agent, agents *[]goap.Agent) *WorldView {
	newAgentsButtons := make([]AgentButton, len(*agents))
	for i, agent := range *agents {
		newAgentsButtons[i] = *NewAgentButton(theme, &agent)
	}

	return &WorldView{
		theme:         *theme,
		world:         *world,
		worldButton:   *NewAgentButton(theme, world),
		agents:        *agents,
		agentsButtons: newAgentsButtons,
	}
}

func (worldView *WorldView) Layout(context layout.Context) layout.Dimensions {
	if worldView.worldButton.Clicked {
		worldView.Selected = worldView.world
	}

	for i, agentsButton := range worldView.agentsButtons {
		if agentsButton.Clicked {
			worldView.Selected = worldView.agents[i]
		}
	}

	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(&worldView.theme, unit.Sp(float32(35)), "World:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(worldView.worldButton.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(&worldView.theme, unit.Sp(float32(35)), "Agents:").Layout),
	}
	for i, _ := range worldView.agentsButtons {
		agentLayouts = append(agentLayouts, layout.Rigid(worldView.agentsButtons[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
