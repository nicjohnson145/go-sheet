package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) lootTab() fyne.CanvasObject {
	contents := make([]*widget.AccordionItem, 0, len(s.character.Loot))
	for _, loot := range s.character.Loot {
		contents = append(contents, s.lootAccordionItem(*loot))
	}

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewAccordion(contents...),
	)
}

func (s *sheet) lootAccordionItem(item CountableItem) *widget.AccordionItem {
	var title string
	if item.Count == 0 {
		title = item.Name
	} else {
		title = fmt.Sprintf("%v (%v)", item.Name, item.Count)
	}

	desc := widget.NewLabel(item.Desc)
	desc.Wrapping = fyne.TextWrapWord
	return widget.NewAccordionItem(title, desc)
}
