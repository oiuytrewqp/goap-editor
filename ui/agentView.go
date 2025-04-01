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
	beliefButtons  *[]BeliefButton
	SelectedBelief int
	goalButtons    *[]GoalButton
	SelectedGoal   *goap.Goal
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
		if len(agentView.agent.Beliefs) == 0 {
			agentView.beliefButtons = &[]BeliefButton{}
		} else {
			beliefButtons := make([]BeliefButton, len(agentView.agent.Beliefs))
			i := 0
			for belief, value := range agentView.agent.Beliefs {
				beliefButtons[i] = *NewBeliefButton(agentView.theme, belief, value)
				i++
			}

			agentView.beliefButtons = &beliefButtons
		}

		if len(agentView.agent.Goals) == 0 {
			agentView.goalButtons = &[]GoalButton{}
		} else {
			goalButtons := make([]GoalButton, len(agentView.agent.Goals))
			for i, goal := range agentView.agent.Goals {
				goalButtons[i] = *NewGoalButton(agentView.theme, &goal)
			}

			agentView.goalButtons = &goalButtons
		}

		if len(agentView.agent.Actions) == 0 {
			agentView.actionButtons = &[]ActionButton{}
		} else {
			actionButtons := make([]ActionButton, len(agentView.agent.Actions))
			for i, action := range agentView.agent.Actions {
				actionButtons[i] = *NewActionButton(agentView.theme, &action)
			}

			agentView.actionButtons = &actionButtons
		}
	}

	for i, beliefButton := range *agentView.beliefButtons {
		if beliefButton.Clicked {
			agentView.SelectedBelief = i
		}
	}

	for i, goalButton := range *agentView.goalButtons {
		if goalButton.Clicked {
			agentView.SelectedGoal = &agentView.agent.Goals[i]
		}
	}

	for i, actionButton := range *agentView.actionButtons {
		if actionButton.Clicked {
			agentView.SelectedAction = &agentView.agent.Actions[i]
		}
	}

	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Agent:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), strconv.Itoa(agentView.agent.Id)).Layout),
		layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), agentView.agent.Name).Layout),
	}

	agentLayouts = append(agentLayouts, layout.Rigid(layout.Spacer{Width: 20}.Layout))
	agentLayouts = append(agentLayouts, layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Beliefs:").Layout))
	if agentView.beliefButtons != nil {
		for i, _ := range *agentView.beliefButtons {
			agentLayouts = append(agentLayouts, layout.Rigid((*agentView.beliefButtons)[i].Layout))
		}
	}

	agentLayouts = append(agentLayouts, layout.Rigid(layout.Spacer{Width: 20}.Layout))
	agentLayouts = append(agentLayouts, layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Goals:").Layout))
	if agentView.goalButtons != nil {
		for i, _ := range *agentView.goalButtons {
			agentLayouts = append(agentLayouts, layout.Rigid((*agentView.goalButtons)[i].Layout))
		}
	}

	agentLayouts = append(agentLayouts, layout.Rigid(layout.Spacer{Width: 20}.Layout))
	agentLayouts = append(agentLayouts, layout.Rigid(material.Label(agentView.theme, unit.Sp(float32(35)), "Actions:").Layout))
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
		agentView.beliefButtons = nil
		agentView.goalButtons = nil
		agentView.actionButtons = nil
	}
}
