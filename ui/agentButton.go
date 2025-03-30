package ui

import (
	"oiuytrewqp/goap-editor/goap"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type AgentButton struct {
	theme   material.Theme
	agent   goap.Agent
	button  widget.Clickable
	Clicked bool
}

func NewAgentButton(theme *material.Theme, agent *goap.Agent) *AgentButton {
	return &AgentButton{
		theme:   *theme,
		agent:   *agent,
		Clicked: false,
	}
}

func (agentButton *AgentButton) Layout(context layout.Context) layout.Dimensions {
	agentButton.Clicked = agentButton.button.Clicked(context)

	return material.Button(&agentButton.theme, &agentButton.button, agentButton.agent.Name).Layout(context)
}
