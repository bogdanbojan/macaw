//go:generate fyne bundle -append -o bundled.go Icon.png

package gui

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	content *widget.Entry
	//	list    *widget.List
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

	hello := widget.NewLabel("Hello!")
	g.win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Choose .txt file", func() {
			g.openFile()
		}),
	))

	w.Resize(fyne.NewSize(500, 320))
	g.win.ShowAndRun()
}

//	func (g *gui) loadUI() fyne.CanvasObject {
//		g.content = widget.NewMultiLineEntry()
//		g.content.SetText(u.placeholderContent())
//
//		visible := g.notes.notes()
//		if len(visible) > 0 {
//			g.setNote(visible[0])
//			g.list.Select(0)
//		}
//
//		bar := widget.NewToolbar(
//			widget.NewToolbarAction(theme.ContentAddIcon(), func() {
//				g.addNote()
//			}),
//			widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
//				g.removeCurrentNote()
//			}),
//		)
//
//		side := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil),
//			bar, container.NewVScroll(g.list))
//
//		return newAdaptiveSplit(side, g.content)
//	}
//
//	func newAdaptiveSplit(left, right fyne.CanvasObject) *fyne.Container {
//		split := container.NewHSplit(left, right)
//		split.Offset = 0.33
//		return container.New(&adaptiveLayout{split: split}, split)
//	}

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
