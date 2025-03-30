package ui

import (
	"oiuytrewqp/goap-editor/goap"
	"strconv"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type AgentView struct {
	theme         material.Theme
	Agent         *goap.Agent
	actionButtons []ActionButton
}

func NewAgentView(theme *material.Theme) *AgentView {
	return &AgentView{
		theme: *theme,
	}
}

func (agentView *AgentView) Layout(context layout.Context) layout.Dimensions {
	if agentView.Agent == nil {
		return material.Label(&agentView.theme, unit.Sp(float32(35)), "Agent:").Layout(context)
	}

	if len(agentView.actionButtons) == 0 {
		agentView.actionButtons = make([]ActionButton, len(agentView.Agent.Actions))
		for i, _ := range agentView.Agent.Actions {
			agentView.actionButtons[i] = *NewActionButton(&agentView.theme, &agentView.Agent.Actions[i])
		}
	}

	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(&agentView.theme, unit.Sp(float32(35)), "Agent:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(&agentView.theme, unit.Sp(float32(35)), strconv.Itoa(agentView.Agent.Id)).Layout),
		layout.Rigid(material.Label(&agentView.theme, unit.Sp(float32(35)), agentView.Agent.Name).Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(&agentView.theme, unit.Sp(float32(35)), "Beliefs:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(&agentView.theme, unit.Sp(float32(35)), "Goals:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(&agentView.theme, unit.Sp(float32(35)), "Actions:").Layout),
	}

	for i, _ := range agentView.actionButtons {
		agentLayouts = append(agentLayouts, layout.Rigid(agentView.actionButtons[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
