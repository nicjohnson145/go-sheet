package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) weaponsTab() fyne.CanvasObject {
	items := make([]*widget.AccordionItem, 0, len(s.character.Weapons))
	for _, wep := range s.character.Weapons {
		items = append(items, s.weaponAccordionItem(*wep))
	}

	return widget.NewAccordion(items...)
}

func (s *sheet) weaponAccordionItem(w Weapon) *widget.AccordionItem {
	return widget.NewAccordionItem(
		fmt.Sprintf(
			"%v (%v)",
			w.Name,
			s.character.modStringForWeapon(w),
		),
		s.weaponAccordionContent(w),
	)
}

func (s *sheet) weaponAccordionContent(w Weapon) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel(w.Desc),
		widget.NewLabel(fmt.Sprintf("Range: %v", w.GetRange())),
		widget.NewLabel(fmt.Sprintf("Damage: %v [%v]", w.Damage.Dice, w.Damage.Type)),
		widget.NewLabel(fmt.Sprintf("Properties: %v", strings.Join(w.Properties, ", "))),
	)
}
