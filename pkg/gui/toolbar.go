package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// constructToolbar initiates the app toolbar with the file search option, the
// sources for the search and an information widget with the repo's URL.
func (g *gui) constructToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar()

	url, _ := url.Parse("https://github.com/bogdanbojan/macaw")
	hyperlink := widget.NewHyperlink("Macaw github repository", url)

	chooseFileToolbar := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		g.openFile()
	})
	settingsToolbar := widget.NewToolbarAction(theme.SettingsIcon(), func() {
		widget.ShowPopUpAtPosition(g.dataFetchContainer, g.win.Canvas(), fyne.NewPos(0, 40))
	})
	infoToolbar := widget.NewToolbarAction(theme.InfoIcon(), func() {
		widget.ShowPopUpAtPosition(hyperlink, g.win.Canvas(), fyne.NewPos(0, 40))

	})
	toolbar.Append(chooseFileToolbar)
	toolbar.Append(settingsToolbar)
	toolbar.Append(infoToolbar)

	g.constructDataFetchContainer()
	return toolbar
}

// constructDataFetchContainer initiates the sliders of the search options found 
// in the toolbar.
func (g *gui) constructDataFetchContainer() {
	g.localDict.slider = widget.NewSlider(0, 1)
	g.onlineDict.slider = widget.NewSlider(0, 1)
	g.wikipedia.slider = widget.NewSlider(0, 1)

	g.dataFetchContainer = container.NewVBox()
	g.dataFetchContainer.Add(widget.NewLabel("Data fetching options"))
	g.dataFetchContainer.Add(widget.NewSeparator())
	g.dataFetchContainer.Add(container.NewAdaptiveGrid(2, widget.NewLabel("Local dictionary"), g.localDict.slider))
	g.dataFetchContainer.Add(container.NewAdaptiveGrid(2, widget.NewLabel("Online dictionary"), g.onlineDict.slider))
	g.dataFetchContainer.Add(container.NewAdaptiveGrid(2, widget.NewLabel("Wikipedia"), g.wikipedia.slider))
}
