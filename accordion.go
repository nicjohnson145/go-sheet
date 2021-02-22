package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type BasicAccordion struct {
	Title string
	Desc string
}

func (s *sheet) scrolledAccordion(items []BasicAccordion) fyne.CanvasObject {
	return container.NewVScroll(widget.NewAccordion(s.makeAccordionItems(items)...))
}

func (s *sheet) makeAccordionItems(items []BasicAccordion) []*widget.AccordionItem {
	accordions := []*widget.AccordionItem{}
	for _, item := range items {
		lbl := widget.NewLabel(item.Desc)
		lbl.Wrapping = fyne.TextWrapWord
		accordion := widget.NewAccordionItem(
			item.Title,
			lbl,
		)
		accordions = append(accordions, accordion)
	}

	return accordions
}
