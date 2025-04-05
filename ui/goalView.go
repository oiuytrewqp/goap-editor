package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type GoalView struct {
	theme material.Theme
	data  *goap.Goap
}

func NewGoalView(theme *material.Theme, data *goap.Goap) *GoalView {
	return &GoalView{
		theme: *theme,
		data:  data,
	}
}

func (goalView *GoalView) Layout(context layout.Context) layout.Dimensions {
	return material.Label(&goalView.theme, unit.Sp(float32(35)), "Goal").Layout(context)
}

func (goalView *GoalView) SetGoal(goal string, value int) {
}
