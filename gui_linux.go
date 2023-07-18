//go:build linux

package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/hotkey"
)

type gui struct {
	search
	win fyne.Window
}

type search struct {
	entry        *widget.Entry
	button       *widget.Button
	result       *widget.Label
	resultScroll *container.Scroll
}

func ShowGUI() {
	a := app.New()
	resourceIconPng, err := fyne.LoadResourceFromPath("./Icon.png")
	if err != nil {
		fmt.Println(err)
	}
	a.SetIcon(resourceIconPng)
	g := &gui{}
	g.win = a.NewWindow("Macaw")

	fnResize := func() {
		g.win.Resize(fyne.NewSize(
			g.win.Canvas().Size().Width,
			300,
			//		g.win.Content().MinSize().Height,
		))
	}

	g.search.entry = widget.NewEntry()
	g.search.entry.SetPlaceHolder("Enter a word to search for...")

	g.search.button = widget.NewButtonWithIcon("Search", theme.SearchIcon(), nil)

	g.search.result = widget.NewLabel("")
	g.search.result.Wrapping = fyne.TextWrapWord
	g.search.resultScroll = container.NewVScroll(g.search.result)

	// TODO: Implement the copy button later.
	// contentCopyButton := widget.NewButtonWithIcon("Copy to clipboard", theme.ContentCopyIcon(), func() {
	// 	g.win.Clipboard().SetContent(g.search.result.Text)
	// })

	g.search.entry.OnSubmitted = func(string) {
		if len(g.search.entry.Text) == 0 {
			dialog.ShowInformation("Search error", "Empty search", g.win)
			return
		}

		// res, err := SearchWiki(g.search.entry.Text)
		// if err != nil {
		// 	g.search.result.SetText("Word not found")
		// 	return
		// }

		def, err := handleLocalResponse(g.search.entry.Text)
		if err != nil || def == nil {
			g.search.result.SetText("Word not found")
			return
		}

		var res string
		for i, v := range def {
			res += fmt.Sprintf("[%d] %s \n", i, v)
		}

		g.search.result.SetText(res)
		fnResize()
	}

	g.win.SetContent(container.NewBorder(
		g.search.entry,
		nil, nil, nil,
		g.search.resultScroll,
	))

	go func() {
		// If numlock is on this will not take effect.
		// Windows+Shift+J
		hk := hotkey.New([]hotkey.Modifier{hotkey.ModShift, hotkey.Mod4}, hotkey.KeyJ)
		if err := hk.Register(); err != nil {
			log.Println("Hotkey registration failed")
		}
		// Start listen hotkey event whenever it is ready.
		for range hk.Keydown() {
			g.win.RequestFocus()
			g.win.Canvas().Focus(g.search.entry)
		}
	}()

	fnResize()
	g.win.Resize(fyne.NewSize(500, 150))
	g.win.ShowAndRun()
}

// TODO: Postpone the features below implementation until we solve the keyboard shortcut problem.
// ===========================================================================

// func (g *gui) constructDataFetchContainer() {
// 	// Data fetching options.
// 	wikiLabel := widget.NewLabel("Wikipedia")
// 	wikiSlider := widget.NewSlider(0, 1)
// 	localDictLabel := widget.NewLabel("Local dictionary")
// 	localDictSlider := widget.NewSlider(0, 1)
// 	onlineDictLabel := widget.NewLabel("Online dictionary")
// 	onlineDictSlider := widget.NewSlider(0, 1)
//
// 	dataFetchContainer := container.NewVBox()
// 	dataFetchContainer.Add(widget.NewLabel("Data fetching options"))
// 	dataFetchContainer.Add(widget.NewSeparator())
// 	dataFetchContainer.Add(container.NewAdaptiveGrid(2, wikiLabel, wikiSlider))
// 	dataFetchContainer.Add(container.NewAdaptiveGrid(2, localDictLabel, localDictSlider))
// 	dataFetchContainer.Add(container.NewAdaptiveGrid(2, onlineDictLabel, onlineDictSlider))
//
// 	g.dataFetchContainer = dataFetchContainer
// }

// func (g *gui) constructToolbar() *widget.Toolbar {
// 	// Use container.NewTabItemWithIcon contained in a container.NewAppTabs
// 	// if you want to add additional text to the icon.
// 	toolbar := widget.NewToolbar()
//
// 	url, _ := url.Parse("https://github.com/bogdanbojan/macaw")
// 	hyperlink := widget.NewHyperlink("Macaw github repository", url)
//
// 	chooseFileToolbar := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
// 		g.openFile()
// 	})
// 	settingsToolbar := widget.NewToolbarAction(theme.SettingsIcon(), func() {
// 		widget.ShowPopUpAtPosition(g.dataFetchContainer, g.win.Canvas(), fyne.NewPos(0, 40))
// 	})
// 	infoToolbar := widget.NewToolbarAction(theme.InfoIcon(), func() {
// 		// dialog.ShowInformation("About", "https://github.com/bogdanbojan/macaw", w)
// 		widget.ShowPopUpAtPosition(hyperlink, g.win.Canvas(), fyne.NewPos(0, 40))
//
// 	})
// 	toolbar.Append(chooseFileToolbar)
// 	toolbar.Append(settingsToolbar)
// 	toolbar.Append(infoToolbar)
//
// 	return toolbar
// }
//
// func (g *gui) openFile() {
// 	dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
// 		if err != nil {
// 			dialog.ShowError(err, g.win)
// 			return
// 		}
// 		if r == nil {
// 			return
// 		}
// 		g.uri = r.URI()
// 		g.loadFile(r)
// 	}, g.win)
// }
//
// func (g *gui) loadFile(r fyne.URIReadCloser) {
// 	read, err := storage.Reader(g.uri)
// 	if err != nil {
// 		log.Println("Error opening resource", err)
// 	}
//
// 	defer read.Close()
// 	data, err := io.ReadAll(read)
// 	if err == nil {
// 		log.Println("Error reading data", err)
// 	}
//
// 	fmt.Println(string(data))
// }
