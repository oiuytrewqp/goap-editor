package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type BeliefView struct {
	theme  material.Theme
	data   *goap.Goap
	belief string
	value  int
}

func NewBeliefView(theme *material.Theme, data *goap.Goap) *BeliefView {
	return &BeliefView{
		theme: *theme,
		data:  data,
	}
}

func (beliefView *BeliefView) Layout(context layout.Context) layout.Dimensions {
	return material.Label(&beliefView.theme, unit.Sp(float32(35)), "Belief").Layout(context)
}

func (beliefView *BeliefView) SetBelief(belief string, value int) {
}
