package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) cardGridFromStrings(values []string) fyne.CanvasObject {
	items := make([]fyne.CanvasObject, 0, len(values))
	for _, val := range values {
		items = append(items, widget.NewCard("", "", widget.NewLabel(val)))
	}

	size := s.minCardSize(items)

	return fyne.NewContainerWithLayout(
		layout.NewGridWrapLayout(size),
		items...,
	)
}

func (s *sheet) minCardSize(items []fyne.CanvasObject) fyne.Size {
	var minwidth float32 = 0
	for _, a := range items {
		size := a.MinSize().Width
		if size > minwidth {
			minwidth = size
		}
	}

	return fyne.NewSize(minwidth, items[0].MinSize().Height)
}
