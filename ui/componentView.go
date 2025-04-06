package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ComponentView struct {
	theme      *material.Theme
	data       *goap.Goap
	AgentList  *AgentList
	BeliefList *BeliefList
	GoalList   *GoalList
	ActionList *ActionList
}

func NewComponentView(theme *material.Theme, data *goap.Goap) *ComponentView {
	return &ComponentView{
		theme:      theme,
		data:       data,
		AgentList:  NewAgentList(theme, data),
		BeliefList: NewBeliefList(theme, data),
		GoalList:   NewGoalList(theme, data),
		ActionList: NewActionList(theme, data),
	}
}

func (componentView *ComponentView) Layout(context layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(context,
		layout.Rigid(componentView.AgentList.Layout),
		layout.Rigid(componentView.BeliefList.Layout),
		layout.Rigid(componentView.GoalList.Layout),
		layout.Rigid(componentView.ActionList.Layout),
	)
}
