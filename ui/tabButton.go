package ui

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type TabButton struct {
	theme   material.Theme
	text    string
	button  widget.Clickable
	Clicked bool
}

func NewTabButton(theme *material.Theme, text string) *TabButton {
	return &TabButton{
		theme:   *theme,
		text:    text,
		Clicked: false,
	}
}

func (tabButton *TabButton) Layout(context layout.Context) layout.Dimensions {
	tabButton.Clicked = tabButton.button.Clicked(context)

	return material.Button(&tabButton.theme, &tabButton.button, tabButton.text).Layout(context)
}
