package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type MainView struct {
	TabView       *TabView
	ComponentView *ComponentView
	AssembleView  *AssembleView
}

func NewMainView(theme *material.Theme, data *goap.Goap) *MainView {
	return &MainView{
		TabView:       NewTabView(theme),
		ComponentView: NewComponentView(theme, data),
		AssembleView:  NewAssembleView(theme, data),
	}
}

func (mainView *MainView) Layout(context layout.Context) layout.Dimensions {
	mainViewLayouts := []layout.FlexChild{
		layout.Rigid(mainView.TabView.Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
	}

	if selected == 0 {
		mainViewLayouts = append(mainViewLayouts, layout.Rigid(mainView.ComponentView.Layout))
	}

	if selected == 1 {
		mainViewLayouts = append(mainViewLayouts, layout.Rigid(mainView.AssembleView.Layout))
	}

	if selected == 2 {
		mainViewLayouts = append(mainViewLayouts, layout.Rigid(mainView.AssembleView.Layout))
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, mainViewLayouts...)
}
