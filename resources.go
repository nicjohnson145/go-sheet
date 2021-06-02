package main

import (
	"fmt"
	"strconv"

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
		resources = append(resources, s.resourceRow(*resource, idx))
	}

	return container.NewVScroll(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			resources...,
		),
	)
}

func (s *sheet) resourceRow(resource Resource, idx int) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		widget.NewButton(
			resource.Name,
			func() { s.onResourceTap(resource, idx) },
		),
		layout.NewSpacer(),
		widget.NewLabel(fmt.Sprintf("%v / %v", resource.Current, resource.Max)),
	)
}


func (s *sheet) onResourceTap(resource Resource, idx int) {
	e := widget.NewEntry()
	modal := widget.NewModalPopUp(nil, s.window.Canvas())
	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel(fmt.Sprintf("Adjust %v", resource.Name)),
		e,
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				widget.NewButton(
					"Remove",
					func() {
						val, err := strconv.Atoi(e.Text)
						if err != nil {
							fmt.Println(err)
							return
						}
						s.character.Resources[idx].Current -= val
						s.writeReadCharacter()
						modal.Hide()
					},
				),
				widget.NewButton(
					"Add",
					func() {
						val, err := strconv.Atoi(e.Text)
						if err != nil {
							fmt.Println(err)
							return
						}
						s.character.Resources[idx].Current += val
						s.writeReadCharacter()
						modal.Hide()
					},
				),
			),
		),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			widget.NewButton(
				"Cancel",
				func() {
					modal.Hide()
				},
			),
		),
	)
	modal.Content = content
	modal.Show()
}
