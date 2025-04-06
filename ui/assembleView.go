package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type AssembleView struct {
	WorldView  *WorldView
	AgentView  *AgentView
	BeliefView *BeliefView
	GoalView   *GoalView
	ActionView *ActionView
}

func NewAssembleView(theme *material.Theme, data *goap.Goap) *AssembleView {
	return &AssembleView{
		WorldView:  NewWorldView(theme, data),
		AgentView:  NewAgentView(theme, data),
		BeliefView: NewBeliefView(theme, data),
		GoalView:   NewGoalView(theme, data),
		ActionView: NewActionView(theme, data),
	}
}

func (assembleView *AssembleView) Layout(context layout.Context) layout.Dimensions {
	assembleView.AgentView.SetAgent(assembleView.WorldView.SelectedAgent)
	//assembleView.BeliefView.SetBelief(assembleView.AgentView.SelectedBelief)
	//assembleView.GoalView.SetGoal(assembleView.AgentView.SelectedGoal)
	assembleView.ActionView.SetAction(assembleView.AgentView.SelectedAction)

	return layout.Flex{}.Layout(context,
		layout.Rigid(assembleView.WorldView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(assembleView.AgentView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(assembleView.BeliefView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(assembleView.GoalView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(assembleView.ActionView.Layout),
	)
}
