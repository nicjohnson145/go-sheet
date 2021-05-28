package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) resourceTab() fyne.CanvasObject {
	resources := []fyne.CanvasObject{}

	for idx, resource := range s.character.Resources {
		if idx > 0 {
			resources = append(resources, widget.NewSeparator())
		}
		resources = append(resources, s.resourceRow(*resource))
	}

	return container.NewVScroll(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			resources...,
		),
	)
}

func (s *sheet) resourceRow(resource Resource) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		widget.NewLabel(resource.Name),
		layout.NewSpacer(),
		widget.NewLabel(fmt.Sprintf("%v / %v", resource.Current, resource.Max)),
	)
}
