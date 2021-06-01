package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
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
			func() {
				e := widget.NewEntry()
				dialog.ShowCustomConfirm(
					fmt.Sprintf("Adjust %v", resource.Name),
					"Add",
					"Remove",
					e,
					func(b bool) {
						val, err := strconv.Atoi(e.Text)
						if err != nil {
							fmt.Println(err)
							return
						}
						if b {
							s.character.Resources[idx].Current += val
							fmt.Println(fmt.Sprintf("Adjusting %v up %v", resource.Name, val))
						} else {
							s.character.Resources[idx].Current -= val
							fmt.Println(fmt.Sprintf("Adjusting %v down %v", resource.Name, val))
						}
						s.writeReadCharacter()
					},
					s.window,
				)
			},
		),
		layout.NewSpacer(),
		widget.NewLabel(fmt.Sprintf("%v / %v", resource.Current, resource.Max)),
	)
}
