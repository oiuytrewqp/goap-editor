package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type WorldView struct {
	theme         *material.Theme
	data          *goap.Goap
	world         *goap.Agent
	worldButton   *AgentButton
	agents        *[]goap.Agent
	agentsButtons *[]AgentButton
	SelectedAgent int
}

func NewWorldView(theme *material.Theme, data *goap.Goap) *WorldView {
	world := data.Agents[data.World]

	worldAgents := make([]goap.Agent, len(data.WorldAgents))
	for i, agent := range data.WorldAgents {
		worldAgents[i] = data.Agents[agent]
	}

	return &WorldView{
		theme:       theme,
		data:        data,
		world:       &world,
		worldButton: NewAgentButton(theme, &world),
		agents:      &worldAgents,
	}
}

func (worldView *WorldView) Layout(context layout.Context) layout.Dimensions {
	if worldView.agentsButtons == nil {
		agentsButtons := make([]AgentButton, len(*worldView.agents))
		for i, agent := range *worldView.agents {
			agentsButtons[i] = *NewAgentButton(worldView.theme, &agent)
		}

		worldView.agentsButtons = &agentsButtons
	}

	if worldView.worldButton.Clicked {
		worldView.SelectedAgent = worldView.world.Id
	}

	for i, agentsButton := range *worldView.agentsButtons {
		if agentsButton.Clicked {
			worldView.SelectedAgent = (*worldView.agents)[i].Id
		}
	}

	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(worldView.theme, unit.Sp(float32(35)), "World:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(worldView.worldButton.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(worldView.theme, unit.Sp(float32(35)), "Agents:").Layout),
	}
	for i, _ := range *worldView.agentsButtons {
		agentLayouts = append(agentLayouts, layout.Rigid((*worldView.agentsButtons)[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
