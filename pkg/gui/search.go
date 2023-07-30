package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bogdanbojan/macaw/pkg/search"
)

const (
	LOCAL  = "LOCAL"
	ONLINE = "ONLINE"
	WIKI   = "WIKI"
)

func (g *gui) initSearchWidgets() {
	g.input.entry = widget.NewEntry()
	g.input.entry.SetPlaceHolder("Enter a word to input for...")

	g.input.button = widget.NewButtonWithIcon("Search", theme.SearchIcon(), nil)

	g.input.localDict.result = widget.NewLabel("")
	g.input.localDict.result.Wrapping = fyne.TextWrapWord
	g.input.localDict.resultScroll = container.NewVScroll(g.input.localDict.result)

	g.input.onlineDict.result = widget.NewLabel("")
	g.input.onlineDict.result.Wrapping = fyne.TextWrapWord
	g.input.onlineDict.resultScroll = container.NewVScroll(g.input.onlineDict.result)

	g.input.wikipedia.result = widget.NewLabel("")
	g.input.wikipedia.result.Wrapping = fyne.TextWrapWord
	g.input.wikipedia.resultScroll = container.NewVScroll(g.input.wikipedia.result)
}

// TODO: Atm, redundant method receiver implementation of the Searcher interface.
func (g *gui) searchWord(string) {
	if len(g.input.entry.Text) == 0 {
		dialog.ShowInformation("Search error", "Empty input", g.win)
		return
	}

	if g.localDict.slider.Value == 1 {
		ld := search.LocalDictionary{}
		g.outputResult(LOCAL, ld.Definition)
	}

	if g.onlineDict.slider.Value == 1 {
		od := search.OnlineDictionary{}
		g.outputResult(ONLINE, od.Definition)
	}

	if g.wikipedia.slider.Value == 1 {
		ws := search.WikipediaSummary{}
		g.outputResult(WIKI, ws.Definition)
	}

	g.tabs.Show()
	g.winResize()
}

func (g *gui) searchWords(ww []string) {
	if g.localDict.slider.Value == 1 {
		ld := search.LocalDictionary{}
		g.outputResults(LOCAL, ld.Definitions, ww)
	}

	if g.onlineDict.slider.Value == 1 {
		od := search.OnlineDictionary{}
		g.outputResults(ONLINE, od.Definitions, ww)
	}

	if g.wikipedia.slider.Value == 1 {
		ws := search.WikipediaSummary{}
		g.outputResults(WIKI, ws.Definitions, ww)
	}

	g.tabs.Show()
	g.winResize()
}

type searchFuncWord func(word string) (string, error)

func (g *gui) outputResult(source string, sf searchFuncWord) {
	res, err := sf(g.input.entry.Text)

	switch source {
	case LOCAL:
		if err != nil {
			g.input.localDict.result.SetText("Word not found")
			return
		}
		g.input.localDict.result.SetText(res)

	case ONLINE:
		if err != nil {
			g.input.onlineDict.result.SetText("Word not found")
			return
		}
		g.input.onlineDict.result.SetText(res)

	case WIKI:
		if err != nil {
			g.input.wikipedia.result.SetText("Word not found")
			return
		}
		g.input.wikipedia.result.SetText(res)
	}
}

type searchFuncWords func(word []string) (string, error)

func (g *gui) outputResults(source string, sf searchFuncWords, ww []string) {
	res, _ := sf(ww)

	switch source {
	case LOCAL:
		g.input.localDict.result.SetText(res)

	case ONLINE:
		g.input.onlineDict.result.SetText(res)

	case WIKI:
		g.input.wikipedia.result.SetText(res)
	}
}
