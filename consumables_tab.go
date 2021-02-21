package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) consumablesTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from consumables tab")
}
