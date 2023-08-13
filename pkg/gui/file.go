package gui

import (
	"io"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

// openFile handles the dialog in which you select your text file for processing.
// It is a controller function which makes use of helper functions to load the
// actual file on your OS and then search the definitions for it.
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

		dataSlice := strings.Split(string(data), "\n")
		g.searchWords(dataSlice)

	}, g.win)
}

// loadFile reads the bytes from your text file.
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
