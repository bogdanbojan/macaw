//go:generate fyne bundle -append -o bundled.go Icon.png

package main

import (
	"fmt"
	"io"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	//	content *widget.Entry
	uri fyne.URI

	win fyne.Window
}

func newGUI(w fyne.Window) *gui {
	return &gui{win: w}
}

func ShowGUI() {
	a := app.New()
	resourceIconPng, err := fyne.LoadResourceFromPath("./Icon.png")
	if err != nil {
		fmt.Println(err)
	}
	a.SetIcon(resourceIconPng)
	w := a.NewWindow("Macaw")
	g := newGUI(w)

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter a word to search for...")

	searchButton := widget.NewButtonWithIcon("Search", theme.SearchIcon(), nil)

	contentCopyButton := widget.NewButtonWithIcon("Copy to clipboard", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(input.Text)
	})
	responseContainer := container.NewVBox()
	responseBox := widget.NewLabel("")
	responseContainer.Add(responseBox)
	responseContainer.Add(contentCopyButton)
	responseContainer.Hidden = true

	// TODOL: Duplicate code. Simplify this later.
	input.OnSubmitted = func(string) {
		if len(input.Text) == 0 {
			dialog.ShowInformation("Search error", "Search entry was empty", w)
			return
		}

		responseContainer.Hidden = false
		responseBox.Text = fmt.Sprintf("You wrote: %s", input.Text)
		responseBox.Refresh()
	}

	searchButton.OnTapped = func() {
		if len(input.Text) == 0 {
			dialog.ShowInformation("Search error", "Search entry was empty", w)
			return
		}

		responseContainer.Hidden = false
		responseBox.Text = fmt.Sprintf("You wrote: %s", input.Text)
		responseBox.Refresh()
		// input.SetText("")
		//	responseBox.Content = container.NewScroll(widget.NewLabel("Test content"))
	}

	// Data fetching options.
	wikiLabel := widget.NewLabel("Wikipedia")
	wikiSlider := widget.NewSlider(0, 1)
	localDictLabel := widget.NewLabel("Local dictionary")
	localDictSlider := widget.NewSlider(0, 1)
	onlineDictLabel := widget.NewLabel("Online dictionary")
	onlineDictSlider := widget.NewSlider(0, 1)

	dataFetchContainer := container.NewVBox()
	dataFetchContainer.Add(widget.NewLabel("Data fetching options"))
	dataFetchContainer.Add(widget.NewSeparator())
	dataFetchContainer.Add(container.NewAdaptiveGrid(2, wikiLabel, wikiSlider))
	dataFetchContainer.Add(container.NewAdaptiveGrid(2, localDictLabel, localDictSlider))
	dataFetchContainer.Add(container.NewAdaptiveGrid(2, onlineDictLabel, onlineDictSlider))

	// Use container.NewTabItemWithIcon contained in a container.NewAppTabs
	// if you want to add additional text to the icon.
	toolbar := widget.NewToolbar()

	url, _ := url.Parse("https://github.com/bogdanbojan/macaw")
	hyperlink := widget.NewHyperlink("Macaw github repository", url)

	chooseFileToolbar := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		g.openFile()
	})
	settingsToolbar := widget.NewToolbarAction(theme.SettingsIcon(), func() {
		widget.ShowPopUpAtPosition(dataFetchContainer, w.Canvas(), fyne.NewPos(0, 40))
	})
	infoToolbar := widget.NewToolbarAction(theme.InfoIcon(), func() {
		// dialog.ShowInformation("About", "https://github.com/bogdanbojan/macaw", w)
		widget.ShowPopUpAtPosition(hyperlink, w.Canvas(), fyne.NewPos(0, 40))

	})
	toolbar.Append(chooseFileToolbar)
	toolbar.Append(settingsToolbar)
	toolbar.Append(infoToolbar)

	g.win.SetContent(container.NewVBox(
		toolbar,
		input,
		searchButton,
		responseContainer,
	))

	g.win.Resize(fyne.NewSize(500, 200))
	g.win.ShowAndRun()
}

func (g *gui) openFile() {
	dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, g.win)
			return
		}
		if r == nil {
			return
		}
		g.uri = r.URI()
		g.loadFile(r)
	}, g.win)
}

func (g *gui) loadFile(r fyne.URIReadCloser) {
	read, err := storage.Reader(g.uri)
	if err != nil {
		log.Println("Error opening resource", err)
	}

	defer read.Close()
	data, err := io.ReadAll(read)
	if err == nil {
		log.Println("Error reading data", err)
	}

	fmt.Println(string(data))
}
