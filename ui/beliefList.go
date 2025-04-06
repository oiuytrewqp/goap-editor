package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type BeliefList struct {
	theme            *material.Theme
	data             *goap.Goap
	beliefComponents *[]Component
}

func NewBeliefList(theme *material.Theme, data *goap.Goap) *BeliefList {
	return &BeliefList{
		theme: theme,
		data:  data,
	}
}

func (beliefList *BeliefList) Layout(context layout.Context) layout.Dimensions {
	if beliefList.beliefComponents == nil {
		beliefComponents := make([]Component, len(beliefList.data.Beliefs))
		i := 0
		for _, agent := range beliefList.data.Beliefs {
			beliefComponents[i] = *NewComponent(beliefList.theme, agent.Name)
			i++
		}

		beliefList.beliefComponents = &beliefComponents
	}
	agentLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(beliefList.theme, unit.Sp(float32(35)), "Beliefs:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
	}
	for i, _ := range *beliefList.beliefComponents {
		agentLayouts = append(agentLayouts, layout.Rigid((*beliefList.beliefComponents)[i].Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, agentLayouts...)
}
