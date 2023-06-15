package frontend

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

	win fyne.Window
}

func newGUI(w fyne.Window) *gui {
	return &gui{win: w}
}

func ShowGUI() {
	a := app.New()
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
			g.content.SetText(string(data))
		}
	}, g.win)
}
