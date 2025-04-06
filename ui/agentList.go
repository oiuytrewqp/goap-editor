package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type AgentList struct {
	theme           *material.Theme
	data            *goap.Goap
	agentComponents *[]Component
}

func NewAgentList(theme *material.Theme, data *goap.Goap) *AgentList {
	return &AgentList{
		theme: theme,
		data:  data,
	}
}

func (agentList *AgentList) Layout(context layout.Context) layout.Dimensions {
	if agentList.agentComponents == nil {
		agentComponents := make([]Component, len(agentList.data.Agents))
		i := 0
		for _, agent := range agentList.data.Agents {
			agentComponents[i] = *NewComponent(agentList.theme, agent.Name)
			i++
		}

		agentList.agentComponents = &agentComponents
	}
	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(agentList.theme, unit.Sp(float32(35)), "Agents:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
	}
	for i, _ := range *agentList.agentComponents {
		agentLayouts = append(agentLayouts, layout.Rigid((*agentList.agentComponents)[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
