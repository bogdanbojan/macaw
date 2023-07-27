package gui

import (
	"errors"
	"fmt"
	"log"

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
		g.outputResult(LOCAL)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResult(ONLINE)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResult(WIKI)
	}

	g.tabs.Show()
	g.winResize()
}

func (g *gui) searchWords(ww []string) {
	if g.localDict.slider.Value == 1 {
		g.outputResults(LOCAL, ww)
	}

	if g.onlineDict.slider.Value == 1 {
		g.outputResults(ONLINE, ww)
	}

	if g.wikipedia.slider.Value == 1 {
		g.outputResults(WIKI, ww)
	}

	g.tabs.Show()
	g.winResize()
}

func (g *gui) outputResult(source string) {
	switch source {
	case LOCAL:
		res, err := g.searchLocalDict(g.search.entry.Text)
		if err != nil {
			g.search.localDict.result.SetText("Word not found")
			return
		}
		g.search.localDict.result.SetText(res)

	case ONLINE:
		res, err := g.searchOnlineDict(g.search.entry.Text)
		if err != nil {
			g.search.onlineDict.result.SetText("Word not found")
			return
		}
		g.search.onlineDict.result.SetText(res)

	case WIKI:
		res, err := g.searchWikipedia(g.search.entry.Text)
		if err != nil {
			g.search.wikipedia.result.SetText("Word not found")
			return
		}
		g.search.wikipedia.result.SetText(res)
	}
}

// TODO: Move for loop repetition into another function.
func (g *gui) outputResults(source string, ww []string) {
	switch source {
	case LOCAL:
		var results []string
		for _, w := range ww {
			res, err := g.searchLocalDict(w)
			if err != nil {
				// TODO: Think about how to handle this. Maybe show a list at
				// the end of words that the app could not find?
				log.Println(err)
				continue
			}
			results = append(results, fmt.Sprint(w+"\n")+res)
		}

		var res string
		for _, v := range results {
			res += fmt.Sprintf(" %s \n", v)
		}

		g.search.localDict.result.SetText(res)

	case ONLINE:
		var results []string
		for _, w := range ww {
			res, err := g.searchOnlineDict(w)
			if err != nil {
				log.Println(err)
				continue
			}
			results = append(results, res)
		}

		var res string
		for _, v := range results {
			res += fmt.Sprintf(" %s \n", v)
		}

		g.search.onlineDict.result.SetText(res)

	case WIKI:
		var results []string
		for _, w := range ww {
			res, err := g.searchWikipedia(w)
			if err != nil {
				log.Println(err)
				continue
			}
			results = append(results, res)
		}

		var res string
		for _, v := range results {
			res += fmt.Sprintf(" %s \n", v)
		}

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
