//go:generate fyne bundle -append -o bundled.go Icon.png

package gui

import (
	"fmt"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	content *widget.Entry
	uri     fyne.URI

	win fyne.Window
}

func newGUI(w fyne.Window) *gui {
	return &gui{win: w}
}

func ShowGUI() {
	a := app.New()
	resourceIconPng, err := fyne.LoadResourceFromPath("./gui/Icon.png")
	if err != nil {
		fmt.Println(err)
	}
	a.SetIcon(resourceIconPng)
	w := a.NewWindow("Macaw")
	g := newGUI(w)

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter a word to search for...")

	searchBox := container.NewVBox(input, widget.NewButtonWithIcon("Search", theme.SearchIcon(), func() {
		log.Println("Content was: ", input.Text)
	}))

	responseText := widget.NewLabel("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	responseText.Wrapping = fyne.TextWrapWord

	responseBox := container.NewScroll(responseText)
	responseBox.SetMinSize(fyne.NewSize(500, 200))

	g.win.SetContent(container.NewVBox(
		// Use container.NewTabItemWithIcon contained in a container.NewAppTabs
		// if you want to add additional text to the icon.
		widget.NewToolbar(
			widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
			widget.NewToolbarAction(theme.InfoIcon(), func() {}),
		),
		widget.NewButton("Choose .txt file", func() {
			g.openFile()
		}),
		searchBox,
		responseBox,
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

		data, err := io.ReadAll(r)
		_ = r.Close()

		if err != nil {
			dialog.ShowError(err, g.win)
		} else {
			g.uri = r.URI()
			g.content.SetText(string(data))
		}
	}, g.win)
}
