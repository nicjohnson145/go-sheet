package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) consumablesTab() fyne.CanvasObject {
	return container.NewVScroll(fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		s.currency(),
		s.itemsAccordion(),
	))
}

func (s *sheet) itemsAccordion() fyne.CanvasObject {
	items := []BasicAccordion{}
	for _, item := range s.character.Consumables {
		items = append(items, BasicAccordion{
			Title: fmt.Sprintf("%v (%v)", item.Name, item.Count),
			Desc:  item.Desc,
		})
	}

	acc := widget.NewAccordion(s.makeAccordionItems(items)...)
	acc.MultiOpen = true
	return acc
}

func (s *sheet) currency() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		s.currencyBox("Copper", &s.character.Currency.Copper),
		s.currencyBox("Silver", &s.character.Currency.Silver),
		s.currencyBox("Electrum", &s.character.Currency.Electrum),
		s.currencyBox("Gold", &s.character.Currency.Gold),
		s.currencyBox("Platinum", &s.character.Currency.Platinum),
	)
}

func (s *sheet) currencyBox(name string, val *int) fyne.CanvasObject {
	return widget.NewButton(
		fmt.Sprintf("%v: %v", name, strconv.Itoa(*val)),
		func() {
			showAddRemoveSetCancelModal(
				s.window.Canvas(),
				AddRemoveSetCancelConfig{
					Label:     fmt.Sprintf("Adjust %v", name),
					Current:   &SetableWrap{Current: val},
					WriteFunc: func() { s.writeReadCharacter() },
				},
			)
		},
	)
}
