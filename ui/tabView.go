package ui

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type TabView struct {
	theme            *material.Theme
	componentsButton *TabButton
	assembleButton   *TabButton
	testButton       *TabButton
}

var selected = 0

func NewTabView(theme *material.Theme) *TabView {
	return &TabView{
		theme:            theme,
		componentsButton: NewTabButton(theme, "Components"),
		assembleButton:   NewTabButton(theme, "Assemble"),
		testButton:       NewTabButton(theme, "Test"),
	}
}

func (tabView *TabView) Layout(context layout.Context) layout.Dimensions {
	if tabView.componentsButton.Clicked {
		selected = 0
	}

	if tabView.assembleButton.Clicked {
		selected = 1
	}

	if tabView.testButton.Clicked {
		selected = 2
	}

	return layout.Flex{}.Layout(context,
		layout.Rigid(tabView.componentsButton.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(tabView.assembleButton.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(tabView.testButton.Layout),
	)
}
