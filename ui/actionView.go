package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type ActionView struct {
	theme  material.Theme
	Action goap.Action
}

func NewActionView(theme *material.Theme) *ActionView {
	return &ActionView{
		theme: *theme,
	}
}

func (actionView *ActionView) Layout(context layout.Context) layout.Dimensions {
	return material.Label(&actionView.theme, unit.Sp(float32(35)), "Action").Layout(context)
}
