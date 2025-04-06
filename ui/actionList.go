package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type ActionList struct {
	theme            *material.Theme
	data             *goap.Goap
	actionComponents *[]Component
}

func NewActionList(theme *material.Theme, data *goap.Goap) *ActionList {
	return &ActionList{
		theme: theme,
		data:  data,
	}
}

func (actionList *ActionList) Layout(context layout.Context) layout.Dimensions {
	if actionList.actionComponents == nil {
		actionComponents := make([]Component, len(actionList.data.Actions))
		i := 0
		for _, agent := range actionList.data.Actions {
			actionComponents[i] = *NewComponent(actionList.theme, agent.Name)
			i++
		}

		actionList.actionComponents = &actionComponents
	}
	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(actionList.theme, unit.Sp(float32(35)), "Beliefs:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
	}
	for i, _ := range *actionList.actionComponents {
		agentLayouts = append(agentLayouts, layout.Rigid((*actionList.actionComponents)[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
