package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type BeliefButton struct {
	theme   material.Theme
	belief  goap.Belief
	button  widget.Clickable
	Clicked bool
}

func NewBeliefButton(theme *material.Theme, belief goap.Belief) *BeliefButton {
	return &BeliefButton{
		theme:   *theme,
		belief:  belief,
		Clicked: false,
	}
}

func (beliefButton *BeliefButton) Layout(context layout.Context) layout.Dimensions {
	beliefButton.Clicked = beliefButton.button.Clicked(context)

	return material.Button(&beliefButton.theme, &beliefButton.button, beliefButton.belief.Name).Layout(context)
}

func (beliefButton *BeliefButton) GetId() int {
	return beliefButton.belief.Id
}
