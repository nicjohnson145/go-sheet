package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type sheet struct {
	window fyne.Window
}

func newSheet() *sheet {
	return &sheet{}
}

func (s *sheet) loadUI(app fyne.App) {
	s.window = app.NewWindow("Go-Sheet")
	s.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			widget.NewLabel("Hello World"),
		),
	)
	s.window.Show()
}
