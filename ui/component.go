package ui

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Component struct {
	theme  material.Theme
	Name   string
	button widget.Clickable
}

func NewComponent(theme *material.Theme, name string) *Component {
	return &Component{
		theme: *theme,
		Name:  name,
	}
}

func (component *Component) Layout(context layout.Context) layout.Dimensions {
	return material.Button(&component.theme, &component.button, component.Name).Layout(context)
}
