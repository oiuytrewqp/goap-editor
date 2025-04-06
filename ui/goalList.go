package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type GoalList struct {
	theme          *material.Theme
	data           *goap.Goap
	goalComponents *[]Component
}

func NewGoalList(theme *material.Theme, data *goap.Goap) *GoalList {
	return &GoalList{
		theme: theme,
		data:  data,
	}
}

func (goalList *GoalList) Layout(context layout.Context) layout.Dimensions {
	if goalList.goalComponents == nil {
		goalComponents := make([]Component, len(goalList.data.Beliefs))
		i := 0
		for _, agent := range goalList.data.Beliefs {
			goalComponents[i] = *NewComponent(goalList.theme, agent.Name)
			i++
		}

		goalList.goalComponents = &goalComponents
	}
	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(goalList.theme, unit.Sp(float32(35)), "Beliefs:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
	}
	for i, _ := range *goalList.goalComponents {
		agentLayouts = append(agentLayouts, layout.Rigid((*goalList.goalComponents)[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
