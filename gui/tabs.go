package gui

import "fyne.io/fyne/v2/container"

func (g *gui) initTabContainers() {
	tabLocalDict := container.NewTabItem("Local dictionary", g.search.localDict.resultScroll)
	tabOnlineDict := container.NewTabItem("Online dictionary", g.search.onlineDict.resultScroll)
	tabWiki := container.NewTabItem("Wikipedia", g.search.wikipedia.resultScroll)

	g.tabs = container.NewAppTabs(tabLocalDict, tabOnlineDict, tabWiki)
	g.tabs.DisableItem(tabOnlineDict)
	g.tabs.DisableItem(tabWiki)
	g.tabs.SetTabLocation(container.TabLocationTop)
}
