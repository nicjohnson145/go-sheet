package main

import (
	"fyne.io/fyne/v2"
)

func (s *sheet) equipmentTab() fyne.CanvasObject {
	items := []BasicAccordion{}
	for _, item := range s.character.Equipment {
		items = append(items, BasicAccordion{
			Title: item.Name,
			Desc:  item.Desc,
		})
	}
	return s.scrolledAccordion(items)
}
