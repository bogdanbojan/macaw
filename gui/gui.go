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
	uri fyne.URI

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
	w := a.NewWindow("Loci")
	g := newGUI(w)

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter a word to search for...")

	searchBox := container.NewVBox(input, widget.NewButtonWithIcon("Search", theme.SearchIcon(), func() {
		log.Println("Content was: ", input.Text)
	}))

	hello := widget.NewLabel("Hello!")
	g.win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Choose .txt file", func() {
			g.openFile()
		}),
		searchBox,
	))

	w.Resize(fyne.NewSize(500, 320))
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
