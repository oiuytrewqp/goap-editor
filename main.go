package main

import (
	"log"
	"oiuytrewqp/goap-editor/goap"
	"oiuytrewqp/goap-editor/ui"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const fullscreen = false
const appTitle = "Goal Oriented Action Plan Editor"

var theme *material.Theme
var context layout.Context

var startButton widget.Clickable

var goapData goap.Goap

func main() {
	goapData = goap.NewGoap()

	go func() {
		window := new(app.Window)
		window.Option(app.Title(appTitle))
		if fullscreen {
			window.Option(app.Fullscreen.Option())
		}
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme = material.NewTheme()
	var ops op.Ops

	mainView := ui.NewMainView(theme, &goapData)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			context = app.NewContext(&ops, e)

			mainView.Layout(context)

			e.Frame(context.Ops)
		}
	}
}
