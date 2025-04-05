package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type GoalButton struct {
	theme   material.Theme
	goal    goap.Goal
	button  widget.Clickable
	Clicked bool
}

func NewGoalButton(theme *material.Theme, goal goap.Goal) *GoalButton {
	return &GoalButton{
		theme:   *theme,
		goal:    goal,
		Clicked: false,
	}
}

func (goalButton *GoalButton) Layout(context layout.Context) layout.Dimensions {
	goalButton.Clicked = goalButton.button.Clicked(context)

	return material.Button(&goalButton.theme, &goalButton.button, goalButton.goal.Name).Layout(context)
}

func (goalButton *GoalButton) GetId() int {
	return goalButton.goal.Id
}
