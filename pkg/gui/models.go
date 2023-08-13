package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// gui is the main struct of the package, which holds all the information regarding
// things like tabs, search input, file URI or window settings.
type gui struct {
	input
	searchOptions      map[string]float64
	dataFetchContainer *fyne.Container
	tabs               *container.AppTabs
	URI                fyne.URI
	win                fyne.Window
}

// TODO: Think of a better name than input.
// input contains all relevant search-related input information.
type input struct {
	entry  *widget.Entry
	button *widget.Button
	sources
}

// sources embeds the search-options of the app, along with widgets that deal 
// with outputting the results in the gui.
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
