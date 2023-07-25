//go:build linux

package gui

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
	"github.com/bogdanbojan/macaw/api"
	"golang.design/x/hotkey"
)

type gui struct {
	search
	dataFetchContainer *fyne.Container
	tabs               *container.AppTabs
	URI                fyne.URI
	win                fyne.Window
}

type search struct {
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

	// TODO: Export this intialization somewhere else.
	g.search.localDict.result = widget.NewLabel("")
	g.search.localDict.result.Wrapping = fyne.TextWrapWord
	g.search.localDict.resultScroll = container.NewVScroll(g.search.localDict.result)

	g.search.onlineDict.result = widget.NewLabel("")
	g.search.onlineDict.result.Wrapping = fyne.TextWrapWord
	g.search.onlineDict.resultScroll = container.NewVScroll(g.search.onlineDict.result)

	g.search.wikipedia.result = widget.NewLabel("")
	g.search.wikipedia.result.Wrapping = fyne.TextWrapWord
	g.search.wikipedia.resultScroll = container.NewVScroll(g.search.wikipedia.result)

	tabLocalDict := container.NewTabItem("Local dictionary", g.search.localDict.resultScroll)
	tabOnlineDict := container.NewTabItem("Online dictionary", g.search.onlineDict.resultScroll)
	tabWiki := container.NewTabItem("Wikipedia", g.search.wikipedia.resultScroll)

	g.tabs = container.NewAppTabs(tabLocalDict, tabOnlineDict, tabWiki)
	g.tabs.DisableItem(tabOnlineDict)
	g.tabs.DisableItem(tabWiki)
	g.tabs.SetTabLocation(container.TabLocationTop)

	g.search.entry.OnSubmitted = func(string) {
		if len(g.search.entry.Text) == 0 {
			dialog.ShowInformation("Search error", "Empty search", g.win)
			return
		}

		if g.localDict.slider.Value == 1 {
			g.tabs.EnableItem(g.tabs.Items[0])

			def, err := api.HandleLocalResponse(g.search.entry.Text)
			if err != nil || def == nil {
				g.search.localDict.result.SetText("Word not found")
				return
			}

			var res string
			for i, v := range def {
				res += fmt.Sprintf("[%d] %s \n", i, v)
			}

			g.search.localDict.result.SetText(res)

		}

		if g.onlineDict.slider.Value == 1 {
			g.tabs.EnableItem(g.tabs.Items[1])

			res, err := api.ApiRequest([]string{g.search.entry.Text})
			if err != nil {
				g.search.onlineDict.result.SetText("Word not found")
				return
			}

			g.search.onlineDict.result.SetText(res[0])
		}

		if g.wikipedia.slider.Value == 1 {
			g.tabs.EnableItem(g.tabs.Items[2])

			res, err := api.SearchWiki(g.search.entry.Text)
			if err != nil {
				g.search.wikipedia.result.SetText("Summary for word not found")
				return
			}
			g.search.wikipedia.result.SetText(res)
		}

		g.tabs.Show()
		fnResize()
	}

	toolbar := g.constructToolbar()
	g.constructDataFetchContainer()
	g.localDict.slider.SetValue(1)

	g.win.SetContent(container.NewBorder(
		toolbar,
		g.search.entry, nil, nil,
		g.tabs,
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
	g.listenSliderChange()
	fnResize()
	g.win.Resize(fyne.NewSize(500, 200))
	g.win.ShowAndRun()
}

func (g *gui) listenSliderChange() {
	g.localDict.slider.OnChanged = func(f float64) {
		if f == 0 {
			g.tabs.DisableItem(g.tabs.Items[0])
			return
		}
		g.tabs.EnableItem(g.tabs.Items[0])
		g.tabs.SelectIndex(0)
	}

	g.onlineDict.slider.OnChanged = func(f float64) {
		if f == 0 {
			g.tabs.DisableItem(g.tabs.Items[1])
			return
		}
		g.tabs.SelectIndex(1)
		g.tabs.EnableItem(g.tabs.Items[1])
	}

	g.wikipedia.slider.OnChanged = func(f float64) {
		if f == 0 {
			g.tabs.DisableItem(g.tabs.Items[2])
			return
		}
		g.tabs.SelectIndex(2)
		g.tabs.EnableItem(g.tabs.Items[2])
	}

}

func (g *gui) constructDataFetchContainer() {
	// Data fetching options.
	localDictLabel := widget.NewLabel("Local dictionary")
	g.localDict.slider = widget.NewSlider(0, 1)
	onlineDictLabel := widget.NewLabel("Online dictionary")
	g.onlineDict.slider = widget.NewSlider(0, 1)
	wikiLabel := widget.NewLabel("Wikipedia")
	g.wikipedia.slider = widget.NewSlider(0, 1)

	dataFetchContainer := container.NewVBox()
	dataFetchContainer.Add(widget.NewLabel("Data fetching options"))
	dataFetchContainer.Add(widget.NewSeparator())
	dataFetchContainer.Add(container.NewAdaptiveGrid(2, localDictLabel, g.localDict.slider))
	dataFetchContainer.Add(container.NewAdaptiveGrid(2, onlineDictLabel, g.onlineDict.slider))
	dataFetchContainer.Add(container.NewAdaptiveGrid(2, wikiLabel, g.wikipedia.slider))

	g.dataFetchContainer = dataFetchContainer
}

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

	return toolbar
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
		g.URI = r.URI()
		g.loadFile(r)
	}, g.win)
}

func (g *gui) loadFile(r fyne.URIReadCloser) {
	read, err := storage.Reader(g.URI)
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
