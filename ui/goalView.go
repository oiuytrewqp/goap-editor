package ui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type GoalView struct {
	theme material.Theme
}

func NewGoalView(theme *material.Theme) *GoalView {
	return &GoalView{
		theme: *theme,
	}
}

func (goalView *GoalView) Layout(context layout.Context) layout.Dimensions {
	return material.Label(&goalView.theme, unit.Sp(float32(35)), "Goal").Layout(context)
}

func (goalView *GoalView) SetGoal(goal string, value int) {
}
