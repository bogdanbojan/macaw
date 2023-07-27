package gui

import (
	"io"
	"log"
	"strings"

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
		data, err := g.loadFile(r)
		if err != nil {
			log.Println("Could not read the file: ", err)
			return
		}
        // TODO: Handle exception where a word is actually a phrase like 
        // George Washington. You don't want [George, Washington].
		dataSlice := strings.Split(string(data), "\n")
		g.searchWords(dataSlice)
        log.Println(dataSlice)

	}, g.win)
}

func (g *gui) loadFile(r fyne.URIReadCloser) ([]byte, error) {
	read, err := storage.Reader(g.URI)
	if err != nil {
		return nil, err
	}

	defer read.Close()
	data, err := io.ReadAll(read)
	if err != nil {
		return nil, err
	}

	return data, nil
}
