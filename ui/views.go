package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Views struct {
	WorldView  *WorldView
	AgentView  *AgentView
	BeliefView *BeliefView
	GoalView   *GoalView
	ActionView *ActionView
}

func NewViews(theme *material.Theme, data *goap.Goap) *Views {
	return &Views{
		WorldView:  NewWorldView(theme, data),
		AgentView:  NewAgentView(theme, data),
		BeliefView: NewBeliefView(theme, data),
		GoalView:   NewGoalView(theme, data),
		ActionView: NewActionView(theme, data),
	}
}

func (views *Views) Layout(context layout.Context) layout.Dimensions {
	views.AgentView.SetAgent(views.WorldView.SelectedAgent)
	//views.BeliefView.SetBelief(views.AgentView.SelectedBelief)
	//views.GoalView.SetGoal(views.AgentView.SelectedGoal)
	views.ActionView.SetAction(views.AgentView.SelectedAction)

	return layout.Flex{}.Layout(context,
		layout.Rigid(views.WorldView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(views.AgentView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(views.BeliefView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(views.GoalView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(views.ActionView.Layout),
	)
}
