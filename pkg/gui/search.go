package gui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bogdanbojan/macaw/pkg/search"
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
	g.searchOptions[search.LOCAL] = g.localDict.slider.Value
	g.searchOptions[search.ONLINE] = g.onlineDict.slider.Value
	g.searchOptions[search.WIKI] = g.wikipedia.slider.Value

	if len(g.input.entry.Text) == 0 {
		dialog.ShowInformation("Search error", "Empty input", g.win)
		return
	}

	if g.localDict.slider.Value == 1 {
		g.outputResult(search.LOCAL, ss.Search)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResult(search.ONLINE, ss.Search)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResult(search.WIKI, ss.Search)
	}

	g.tabs.Show()
	g.winResize()
}

func (g *gui) searchWords(ww []string) {
	ss := search.Sources{}

	if g.localDict.slider.Value == 1 {
		g.outputResults(search.LOCAL, ss.Search, ww)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResults(search.ONLINE, ss.Search, ww)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResults(search.WIKI, ss.Search, ww)
	}

	g.tabs.Show()
	g.winResize()
}

type searchFunc func(ctx context.Context, words []string) []search.Definition

func (g *gui) outputResult(source string, sf searchFunc) {
	ctx := context.WithValue(context.Background(), search.ContextKeyOptions, g.searchOptions)

	res := sf(ctx, []string{g.input.entry.Text})

	switch source {
	case search.LOCAL:
		if !res[0].Ok {
			g.input.localDict.result.SetText("Word not found")
			return
		}
		g.input.localDict.result.SetText(res[0].Text)

	case search.ONLINE:
		if !res[0].Ok {
			g.input.onlineDict.result.SetText("Word not found")
			return
		}
		g.input.onlineDict.result.SetText(res[0].Text)

	case search.WIKI:
		if !res[0].Ok {
			g.input.wikipedia.result.SetText("Word not found")
			return
		}
		g.input.wikipedia.result.SetText(res[0].Text)
	}
}

// TODO: Set limit amount on word input?
func (g *gui) outputResults(source string, sf searchFunc, ww []string) {
	ctx := context.WithValue(context.Background(), search.ContextKeyOptions, g.searchOptions)

	res := sf(ctx, ww)

	switch source {
	case search.LOCAL:
		dd, fdd := splitDefinitions(res)
		g.input.localDict.result.SetText(toString(dd, fdd))

	case search.ONLINE:
		dd, fdd := splitDefinitions(res)
		g.input.localDict.result.SetText(toString(dd, fdd))

	case search.WIKI:
		dd, fdd := splitDefinitions(res)
		g.input.localDict.result.SetText(toString(dd, fdd))
	}
}

func splitDefinitions(res []search.Definition) (definitions, failedDefinitions []string) {
	var dd []string
	var fdd []string
	for _, r := range res {
		if !r.Ok {
			fdd = append(fdd, r.Word)
			continue
		}
		dd = append(dd, r.Text)
	}
	return dd, fdd
}

func toString(definitions, failedDefinitions []string) string {
	var stringdd string
	for _, d := range definitions {
		stringdd += fmt.Sprintf(" %s \n", d)
	}

	if len(failedDefinitions) != 0 {
		stringdd += "Could not find the following words: \n"
		for _, fd := range failedDefinitions {
			stringdd += fd + "\n"
		}
	}
	return stringdd
}
