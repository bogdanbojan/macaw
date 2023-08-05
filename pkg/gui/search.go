package gui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bogdanbojan/macaw/pkg/search"
)

var (
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
	ss := search.Sources{}
	g.searchOptions = make(map[string]float64, 3)
	g.searchOptions[LOCAL] = g.localDict.slider.Value
	g.searchOptions[ONLINE] = g.onlineDict.slider.Value
	g.searchOptions[WIKI] = g.wikipedia.slider.Value

	if len(g.input.entry.Text) == 0 {
		dialog.ShowInformation("Search error", "Empty input", g.win)
		return
	}

	if g.localDict.slider.Value == 1 {
		g.outputResult(LOCAL, ss.Search)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResult(ONLINE, ss.Search)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResult(WIKI, ss.Search)
	}

	g.tabs.Show()
	g.winResize()
}

func (g *gui) searchWords(ww []string) {
	ss := search.Sources{}

	if g.localDict.slider.Value == 1 {
		g.outputResults(LOCAL, ss.Search, ww)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResults(ONLINE, ss.Search, ww)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResults(WIKI, ss.Search, ww)
	}

	g.tabs.Show()
	g.winResize()
}

type searchFunc func(ctx context.Context, words []string) (string, error)

func (g *gui) outputResult(source string, sf searchFunc) {
	ctx := context.WithValue(context.Background(), search.ContextKeyOptions, g.searchOptions)

	res, err := sf(ctx, []string{g.input.entry.Text})

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

func (g *gui) outputResults(source string, sf searchFunc, ww []string) {
	ctx := context.WithValue(context.Background(), search.ContextKeyOptions, g.searchOptions)

	res, _ := sf(ctx, ww)

	switch source {
	case LOCAL:
		g.input.localDict.result.SetText(res)

	case ONLINE:
		g.input.onlineDict.result.SetText(res)

	case WIKI:
		g.input.wikipedia.result.SetText(res)
	}
}
