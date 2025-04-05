package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ActionButton struct {
	theme   material.Theme
	action  goap.Action
	button  widget.Clickable
	Clicked bool
}

func NewActionButton(theme *material.Theme, action goap.Action) *ActionButton {
	return &ActionButton{
		theme:   *theme,
		action:  action,
		Clicked: false,
	}
}

func (actionButton *ActionButton) Layout(context layout.Context) layout.Dimensions {
	actionButton.Clicked = actionButton.button.Clicked(context)

	return material.Button(&actionButton.theme, &actionButton.button, actionButton.action.Name).Layout(context)
}

func (actionButton *ActionButton) GetId() int {
	return actionButton.action.Id
}
