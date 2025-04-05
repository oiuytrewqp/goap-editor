package ui

import (
	"oiuytrewqp/goap-editor/goap"
	"strconv"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type ActionView struct {
	theme                *material.Theme
	data                 *goap.Goap
	action               *goap.Action
	prerequisitesButtons *[]StateButton
	outcomesButtons      *[]StateButton
}

func NewActionView(theme *material.Theme, data *goap.Goap) *ActionView {
	return &ActionView{
		theme: theme,
		data:  data,
	}
}

// id, name, prereqs, outcomes
func (actionView *ActionView) Layout(context layout.Context) layout.Dimensions {
	if actionView.action == nil {
		return material.Label(actionView.theme, unit.Sp(float32(35)), "Action:").Layout(context)
	}

	if actionView.prerequisitesButtons == nil {
		if len(actionView.action.Prerequisites) == 0 {
			actionView.prerequisitesButtons = &[]StateButton{}
		} else {
			prerequisitesButtons := make([]StateButton, len(actionView.action.Prerequisites))
			i := 0
			for prerequisite, value := range actionView.action.Prerequisites {
				prerequisitesButtons[i] = *NewStateButton(actionView.theme, actionView.data.Beliefs[prerequisite], value)
				i++
			}

			actionView.prerequisitesButtons = &prerequisitesButtons
		}

		if len(actionView.action.Outcomes) == 0 {
			actionView.outcomesButtons = &[]StateButton{}
		} else {
			outcomesButtons := make([]StateButton, len(actionView.action.Outcomes))
			i := 0
			for outcome, value := range actionView.action.Outcomes {
				outcomesButtons[i] = *NewStateButton(actionView.theme, actionView.data.Beliefs[outcome], value)
				i++
			}

			actionView.outcomesButtons = &outcomesButtons
		}
	}

	actionLayouts := []layout.FlexChild{
		layout.Rigid(material.Label(actionView.theme, unit.Sp(float32(35)), "Action:").Layout),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(material.Label(actionView.theme, unit.Sp(float32(35)), strconv.Itoa(actionView.action.Id)).Layout),
		layout.Rigid(material.Label(actionView.theme, unit.Sp(float32(35)), actionView.action.Name).Layout),
	}

	actionLayouts = append(actionLayouts, layout.Rigid(layout.Spacer{Width: 20}.Layout))
	actionLayouts = append(actionLayouts, layout.Rigid(material.Label(actionView.theme, unit.Sp(float32(35)), "Prerequisites:").Layout))
	if actionView.prerequisitesButtons != nil {
		for i, _ := range *actionView.prerequisitesButtons {
			actionLayouts = append(actionLayouts, layout.Rigid((*actionView.prerequisitesButtons)[i].Layout))
		}
	}

	actionLayouts = append(actionLayouts, layout.Rigid(layout.Spacer{Width: 20}.Layout))
	actionLayouts = append(actionLayouts, layout.Rigid(material.Label(actionView.theme, unit.Sp(float32(35)), "Outcomes:").Layout))
	if actionView.outcomesButtons != nil {
		for i, _ := range *actionView.outcomesButtons {
			actionLayouts = append(actionLayouts, layout.Rigid((*actionView.outcomesButtons)[i].Layout))
		}
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(context, actionLayouts...)
}

func (actionView *ActionView) SetAction(id int) {
	if id == 0 {
		return
	}

	newAction := actionView.data.Actions[id]

	if actionView.action == nil || actionView.action.Id != newAction.Id {
		actionView.action = &newAction
		actionView.prerequisitesButtons = nil
		actionView.outcomesButtons = nil
	}
}
