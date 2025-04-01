package ui

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type BeliefButton struct {
	theme   material.Theme
	belief  string
	value   int
	button  widget.Clickable
	Clicked bool
}

func NewBeliefButton(theme *material.Theme, belief string, value int) *BeliefButton {
	return &BeliefButton{
		theme:   *theme,
		belief:  belief,
		value:   value,
		Clicked: false,
	}
}

func (beliefButton *BeliefButton) Layout(context layout.Context) layout.Dimensions {
	beliefButton.Clicked = beliefButton.button.Clicked(context)

	return material.Button(&beliefButton.theme, &beliefButton.button, beliefButton.belief).Layout(context)
}
