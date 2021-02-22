package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (s *sheet) consumablesTab() fyne.CanvasObject {
	items := []BasicAccordion{}
	for _, item := range s.character.Consumables {
		items = append(items, BasicAccordion{
			Title: fmt.Sprintf("%v (%v)", item.Name, item.Count),
			Desc: item.Desc,
		})
	}

	return s.scrolledAccordion(items)
}
