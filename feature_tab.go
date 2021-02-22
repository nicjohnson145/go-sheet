package main

import (
	"fyne.io/fyne/v2"
)

func (s *sheet) featuresTab() fyne.CanvasObject {
	items := []BasicAccordion{}
	for _, item := range s.character.Features {
		items = append(items, BasicAccordion{
			Title: item.Name,
			Desc: item.Desc,
		})
	}

	return s.scrolledAccordion(items)
}
