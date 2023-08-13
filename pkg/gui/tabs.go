package gui

import "fyne.io/fyne/v2/container"

// initTabContainers initiates the tabs, found in the toolbar, which hold the search options.
func (g *gui) initTabContainers() {
	tabLocalDict := container.NewTabItem("Local dictionary", g.input.localDict.resultScroll)
	tabOnlineDict := container.NewTabItem("Online dictionary", g.input.onlineDict.resultScroll)
	tabWiki := container.NewTabItem("Wikipedia", g.input.wikipedia.resultScroll)

	g.tabs = container.NewAppTabs(tabLocalDict, tabOnlineDict, tabWiki)
	g.tabs.DisableItem(tabOnlineDict)
	g.tabs.DisableItem(tabWiki)
	g.tabs.SetTabLocation(container.TabLocationTop)
}
