package gui

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bogdanbojan/macaw/api"
)

const (
	LOCAL  = "LOCAL"
	ONLINE = "ONLINE"
	WIKI   = "WIKI"
)

func (g *gui) initSearchWidgets() {
	g.search.entry = widget.NewEntry()
	g.search.entry.SetPlaceHolder("Enter a word to search for...")

	g.search.button = widget.NewButtonWithIcon("Search", theme.SearchIcon(), nil)

	g.search.localDict.result = widget.NewLabel("")
	g.search.localDict.result.Wrapping = fyne.TextWrapWord
	g.search.localDict.resultScroll = container.NewVScroll(g.search.localDict.result)

	g.search.onlineDict.result = widget.NewLabel("")
	g.search.onlineDict.result.Wrapping = fyne.TextWrapWord
	g.search.onlineDict.resultScroll = container.NewVScroll(g.search.onlineDict.result)

	g.search.wikipedia.result = widget.NewLabel("")
	g.search.wikipedia.result.Wrapping = fyne.TextWrapWord
	g.search.wikipedia.resultScroll = container.NewVScroll(g.search.wikipedia.result)
}

func (g *gui) searchWord(string) {
	if len(g.search.entry.Text) == 0 {
		dialog.ShowInformation("Search error", "Empty search", g.win)
		return
	}

	if g.localDict.slider.Value == 1 {
		g.outputResult(LOCAL, g.searchLocalDict)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResult(ONLINE, g.searchOnlineDict)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResult(WIKI, g.searchWikipedia)
	}

	g.tabs.Show()
	g.winResize()
}

func (g *gui) searchWords(ww []string) {
	if g.localDict.slider.Value == 1 {
		g.outputResults(LOCAL, g.searchLocalDict, ww)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResults(ONLINE, g.searchOnlineDict, ww)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResults(WIKI, g.searchWikipedia, ww)
	}

	g.tabs.Show()
	g.winResize()
}

type searchFunc func(word string) (string, error)

func (g *gui) outputResult(source string, sf searchFunc) {
	res, err := sf(g.search.entry.Text)

	switch source {
	case LOCAL:
		if err != nil {
			g.search.localDict.result.SetText("Word not found")
			return
		}
		g.search.localDict.result.SetText(res)

	case ONLINE:
		if err != nil {
			g.search.onlineDict.result.SetText("Word not found")
			return
		}
		g.search.onlineDict.result.SetText(res)

	case WIKI:
		if err != nil {
			g.search.wikipedia.result.SetText("Word not found")
			return
		}
		g.search.wikipedia.result.SetText(res)
	}
}

func (g *gui) outputResults(source string, sf searchFunc, ww []string) {
	res := func(searchFunc) string {
		var results []string
		var failedResults []string
		for _, w := range ww {
			res, err := sf(w)
			if err != nil {
				failedResults = append(failedResults, w)
				continue
			}
			results = append(results, fmt.Sprint(w+"\n")+res)
		}

		var res string
		for _, v := range results {
			res += fmt.Sprintf(" %s \n", v)
		}

		if len(failedResults) != 0 {
			res += "Could not find the following words: \n"
			for _, v := range failedResults {
				res += v + "\n"
			}
		}

		return res
	}(sf)

	switch source {
	case LOCAL:
		g.search.localDict.result.SetText(res)

	case ONLINE:
		g.search.onlineDict.result.SetText(res)

	case WIKI:
		g.search.wikipedia.result.SetText(res)
	}
}

func (g *gui) searchLocalDict(word string) (string, error) {
	g.tabs.EnableItem(g.tabs.Items[0])

	def, err := api.GetLocalDefinition(word)
	if err != nil || def == nil {
		return "", errors.New("word not found")
	}

	var res string
	for i, v := range def {
		res += fmt.Sprintf("[%d] %s \n", i, v)
	}

	return res, nil
}

func (g *gui) searchOnlineDict(word string) (string, error) {
	g.tabs.EnableItem(g.tabs.Items[1])

	res, err := api.GetOnlineDefinition(word)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (g *gui) searchWikipedia(word string) (string, error) {
	g.tabs.EnableItem(g.tabs.Items[2])

	res, err := api.GetWikipediaSummary(word)
	if err != nil {
		return "", err
	}

	return res, nil
}
