package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) weaponsTab() fyne.CanvasObject {
	items := make([]*widget.AccordionItem, 0, len(s.character.Weapons))
	for _, wep := range s.character.Weapons {
		items = append(items, s.weaponAccordionItem(*wep))
	}

	acc := widget.NewAccordion(items...)
	acc.MultiOpen = true
	return container.NewVScroll(acc)
}

func (s *sheet) weaponAccordionItem(w Weapon) *widget.AccordionItem {
	return widget.NewAccordionItem(
		fmt.Sprintf(
			"%v (%v) [%v%v]",
			w.Name,
			s.character.modStringForWeapon(w),
			w.Damage.Dice,
			s.character.modOnlyModString(s.character.calcMod(s.character.attrForString(w.Attribute)) + w.Damage.AdditionalDamage),
		),
		s.weaponAccordionContent(w),
	)
}

func (s *sheet) weaponAccordionContent(w Weapon) fyne.CanvasObject {
	attrs, size := s.weaponItemAnnotations(w)

	body := []fyne.CanvasObject{}
	if w.Desc != "" {
		desc := widget.NewLabel(w.Desc)
		desc.Wrapping = fyne.TextWrapWord
		body = append(body, desc)
	}

	body = append(body, fyne.NewContainerWithLayout(
		layout.NewGridWrapLayout(size),
		attrs...,
	))

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		body...,
	)
}

func (s *sheet) weaponItemAnnotations(w Weapon) ([]fyne.CanvasObject, fyne.Size) {
	attrs := []fyne.CanvasObject{
		widget.NewCard("", "", widget.NewLabel(fmt.Sprintf("Range: %v", w.GetRange()))),
		widget.NewCard("", "", widget.NewLabel(fmt.Sprintf("Damage: %v [%v]", w.Damage.Dice, w.Damage.Type))),
		widget.NewCard("", "", widget.NewLabel(fmt.Sprintf("Properties: %v", strings.Join(w.Properties, ", ")))),
	}

	size := s.minCardSize(attrs)

	return attrs, size
}
