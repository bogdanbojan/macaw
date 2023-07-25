package gui

import (
	"fmt"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

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
