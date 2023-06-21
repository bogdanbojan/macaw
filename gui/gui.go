//go:generate fyne bundle -append -o bundled.go Icon.png

package gui

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	w := a.NewWindow("Loci")
	g := newGUI(w)

	hello := widget.NewLabel("Hello Fyne!")
	g.win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		widget.NewButton("Open File", func() {
			g.openFile()
		}),
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
