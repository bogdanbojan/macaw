package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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

func (g *gui) constructDataFetchContainer() {
	localDictLabel := widget.NewLabel("Local dictionary")
	g.localDict.slider = widget.NewSlider(0, 1)
	onlineDictLabel := widget.NewLabel("Online dictionary")
	g.onlineDict.slider = widget.NewSlider(0, 1)
	wikiLabel := widget.NewLabel("Wikipedia")
	g.wikipedia.slider = widget.NewSlider(0, 1)

	g.dataFetchContainer = container.NewVBox()
	g.dataFetchContainer.Add(widget.NewLabel("Data fetching options"))
	g.dataFetchContainer.Add(widget.NewSeparator())
	g.dataFetchContainer.Add(container.NewAdaptiveGrid(2, localDictLabel, g.localDict.slider))
	g.dataFetchContainer.Add(container.NewAdaptiveGrid(2, onlineDictLabel, g.onlineDict.slider))
	g.dataFetchContainer.Add(container.NewAdaptiveGrid(2, wikiLabel, g.wikipedia.slider))
}
