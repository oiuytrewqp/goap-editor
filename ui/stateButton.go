package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type StateButton struct {
	theme   material.Theme
	belief  goap.Belief
	button  widget.Clickable
	Clicked bool
}

func NewStateButton(theme *material.Theme, belief goap.Belief, value int) *StateButton {
	return &StateButton{
		theme:   *theme,
		belief:  belief,
		Clicked: false,
	}
}

func (stateButton *StateButton) Layout(context layout.Context) layout.Dimensions {
	stateButton.Clicked = stateButton.button.Clicked(context)

	return material.Button(&stateButton.theme, &stateButton.button, stateButton.belief.Name).Layout(context)
}
