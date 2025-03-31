package ui

import (
	"oiuytrewqp/goap-editor/goap"
	"strconv"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type AgentView struct {
	theme          *material.Theme
	agent          *goap.Agent
	actions        *[]goap.Action
	actionButtons  *[]ActionButton
	SelectedAction *goap.Action
}

func NewAgentView(theme *material.Theme) *AgentView {
	return &AgentView{
		theme: theme,
	}
}

func (agentView *AgentView) Layout(context layout.Context) layout.Dimensions {
	if agentView.agent == nil {
		return material.Label(agentView.theme, unit.Sp(float32(35)), "Agent:").Layout(context)
	}

	if agentView.actionButtons == nil {
		if len(*agentView.actions) == 0 {
			return material.Label(agentView.theme, unit.Sp(float32(35)), "Agent:").Layout(context)
		}

		actionButtons := make([]ActionButton, len(*agentView.actions))
		for i, action := range *agentView.actions {
			actionButtons[i] = *NewActionButton(agentView.theme, &action)
		}

		agentView.actionButtons = &actionButtons
	}

	for i, actionButton := range *agentView.actionButtons {
		if actionButton.Clicked {
			agentView.SelectedAction = &(*agentView.actions)[i]
		}
	}

	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Agent:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), strconv.Itoa(agentView.agent.Id)).Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), agentView.agent.Name).Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Beliefs:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Goals:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Actions:").Layout),
	}

	if agentView.actionButtons != nil {
		for i, _ := range *agentView.actionButtons {
			agentLayouts = append(agentLayouts, layout.Rigid((*agentView.actionButtons)[i].Layout))
		}
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}

func (agentView *AgentView) SetAgent(agent *goap.Agent) {
	if agentView.agent != agent {
		agentView.agent = agent
		if agent != nil {
			agentView.actions = &agent.Actions
		}

		agentView.actionButtons = nil
	}
}
