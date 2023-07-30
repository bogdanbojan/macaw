package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	input
	dataFetchContainer *fyne.Container
	tabs               *container.AppTabs
	URI                fyne.URI
	win                fyne.Window
}

type input struct {
	entry  *widget.Entry
	button *widget.Button
	sources
}

type sources struct {
	localDict
	onlineDict
	wikipedia
}

type localDict struct {
	result       *widget.Label
	resultScroll *container.Scroll
	slider       *widget.Slider
}

type onlineDict struct {
	result       *widget.Label
	resultScroll *container.Scroll
	slider       *widget.Slider
}

type wikipedia struct {
	result       *widget.Label
	resultScroll *container.Scroll
	slider       *widget.Slider
}
