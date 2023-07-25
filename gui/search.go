package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bogdanbojan/macaw/api"
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

// TODO: Fix bug regarding the search not working properly on the online dictionary
// when the local dictionary does not find the word. It has to do with the return
// statement in the conditionals.
func (g *gui) searchSources(string) {
	if len(g.search.entry.Text) == 0 {
		dialog.ShowInformation("Search error", "Empty search", g.win)
		return
	}

	if g.localDict.slider.Value == 1 {
		g.tabs.EnableItem(g.tabs.Items[0])

		def, err := api.HandleLocalResponse(g.search.entry.Text)
		if err != nil || def == nil {
			g.search.localDict.result.SetText("Word not found")
			return
		}

		var res string
		for i, v := range def {
			res += fmt.Sprintf("[%d] %s \n", i, v)
		}

		g.search.localDict.result.SetText(res)

	}

	if g.onlineDict.slider.Value == 1 {
		g.tabs.EnableItem(g.tabs.Items[1])

		res, err := api.ApiRequest([]string{g.search.entry.Text})
		if err != nil {
			g.search.onlineDict.result.SetText("Word not found")
			return
		}
        log.Println(g.search.entry.Text)

		g.search.onlineDict.result.SetText(res[0])
	}

	if g.wikipedia.slider.Value == 1 {
		g.tabs.EnableItem(g.tabs.Items[2])

		res, err := api.SearchWiki(g.search.entry.Text)
		if err != nil {
			g.search.wikipedia.result.SetText("Summary for word not found")
			return
		}
		g.search.wikipedia.result.SetText(res)
	}

	g.tabs.Show()
	g.winResize()
}
