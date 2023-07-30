//go:build darwin

package gui

import (
	"log"

	"golang.design/x/hotkey"
)

func (g *gui) initHotkey() {
	// If numlock is on this will not take effect.
	// Windows+Shift+J
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
