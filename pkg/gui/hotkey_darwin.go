//go:build darwin

package gui

import (
	"log"

	"golang.design/x/hotkey"
)

// initHotkey registers a system wide keyboard shortcut for the application pop-up.
// On darwin systems, the keyboard shortcut is `Option Key + Shift Key + J`.
func (g *gui) initHotkey() {
	// If numlock is on this will not take effect.
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModShift, hotkey.ModOption}, hotkey.KeyJ)
	if err := hk.Register(); err != nil {
		log.Println("Hotkey registration failed")
	}
	// Start listen hotkey event whenever it is ready.
	for range hk.Keydown() {
		g.win.RequestFocus()
		g.win.Canvas().Focus(g.input.entry)
	}
}
