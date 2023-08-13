package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/bogdanbojan/macaw/pkg/gui/assets"
)

// ShowGUI is the main controller function for the GUI. It sets up the widgets,
// toolbar, searchbar and content for the user.
func ShowGUI() {
	a := app.New()
	a.SetIcon(assets.AppIcon)
	g := &gui{}
	g.win = a.NewWindow("Macaw")

	g.initSearchWidgets()
	g.initTabContainers()
	toolbar := g.constructToolbar()

	g.input.entry.OnSubmitted = func(s string) { g.searchWord(s) }
	g.localDict.slider.SetValue(1)
	go g.initHotkey()

	g.win.SetContent(container.NewBorder(
		toolbar,
		g.input.entry, nil, nil,
		g.tabs,
	))

	g.listenSliderChange()
	g.winResize()
	g.win.Resize(fyne.NewSize(500, 200))
	g.win.ShowAndRun()
}

// winResize refreshes the size of the window. It's needed to change the size 
// dynamically when you make a search in the app.
func (g *gui) winResize() {
	g.win.Resize(fyne.NewSize(
		g.win.Canvas().Size().Width,
		300,
		//		g.win.Content().MinSize().Height,
	))
}

// TODO: Add logic to select last activated tab when deactivating a certain tab.
// listenSliderChange listens to any search option changes that the user makes.
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
