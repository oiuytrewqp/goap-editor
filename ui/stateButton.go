package ui

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type StateButton struct {
	theme   material.Theme
	name    string
	button  widget.Clickable
	Clicked bool
}

func NewStateButton(theme *material.Theme, name string, value int) *StateButton {
	return &StateButton{
		theme:   *theme,
		name:    name,
		Clicked: false,
	}
}

func (stateButton *StateButton) Layout(context layout.Context) layout.Dimensions {
	stateButton.Clicked = stateButton.button.Clicked(context)

	return material.Button(&stateButton.theme, &stateButton.button, stateButton.name).Layout(context)
}
