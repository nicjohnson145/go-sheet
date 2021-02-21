package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) featuresTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from feature tab")
}
