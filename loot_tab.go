package main

import (
	"fmt"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) lootTab() fyne.CanvasObject {
	contents := make([]*widget.AccordionItem, 0, len(s.character.Loot))

	allLoot := make([]*CountableItem, len(s.character.Loot), len(s.character.Loot))
	copy(allLoot, s.character.Loot)
	sort.Slice(allLoot, func(i, j int) bool {
		return allLoot[i].Name < allLoot[j].Name
	})

	for _, loot := range allLoot {
		contents = append(contents, s.lootAccordionItem(*loot))
	}

	return container.NewVScroll(widget.NewAccordion(contents...))
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
