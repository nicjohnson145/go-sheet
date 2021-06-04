package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) spellTab() fyne.CanvasObject {
	sections := []fyne.CanvasObject{}
	levels := []string{
		Cantrips,
		Level1,
		Level2,
		Level3,
		Level4,
		Level5,
		Level6,
		Level7,
		Level8,
		Level9,
	}

	for _, name := range levels {
		if _, ok := s.character.Spells[name]; ok {
			sections = append(sections, s.spellSection(name))
		}
	}

	return container.NewVScroll(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			sections...,
		),
	)
}

func (s *sheet) spellSection(name string) fyne.CanvasObject {
	var titleRow []fyne.CanvasObject
	if name == Cantrips {
		titleRow = []fyne.CanvasObject{
			widget.NewLabel(name),
			layout.NewSpacer(),
		}
	} else {
		titleRow = []fyne.CanvasObject{
			widget.NewButton(
				name,
				func() {
					showAddRemoveSetCancelModal(
						s.window.Canvas(),
						AddRemoveSetCancelConfig{
							Label: "Adjust Slots",
							Current: s.character.Spells[name],
							WriteFunc: func() { s.writeReadCharacter() },
						},
					)
				},
			),
			layout.NewSpacer(),
		}
		titleRow = s.addSlotInfo(titleRow, *s.character.Spells[name])
	}
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			titleRow...,
		),
		s.spellAccordion(s.character.Spells[name].Spells),
	)
}

func (s *sheet) addSlotInfo(row []fyne.CanvasObject, section SpellSection) []fyne.CanvasObject {
	for i := 1; i <= section.MaxSlots; i++ {
		if i <= section.Slots {
			row = append(row, widget.NewIcon(theme.RadioButtonCheckedIcon()))
		} else {
			row = append(row, widget.NewIcon(theme.RadioButtonIcon()))
		}
	}

	return row
}

func (s *sheet) spellAccordion(spells []*Spell) fyne.CanvasObject {
	items := []*widget.AccordionItem{}
	for _, spell := range spells {
		items = append(items, widget.NewAccordionItem(spell.Name, s.spellCard(*spell)))
	}
	acc := widget.NewAccordion(items...)
	acc.MultiOpen = true
	return acc
}

func (s *sheet) spellCard(spell Spell) fyne.CanvasObject {
	cards := []string{
		fmt.Sprintf("Range: %v", spell.Range),
		fmt.Sprintf("Duration: %v", spell.Duration),
		fmt.Sprintf("Concentration: %v", s.boolToStr(spell.Concentration)),
		fmt.Sprintf("Ritual: %v", s.boolToStr(spell.Ritual)),
		fmt.Sprintf("Components: %v", strings.Join(spell.Components, ", ")),
	}

	grid := s.cardGridFromStrings(cards)

	desc := widget.NewLabel(spell.Desc)
	desc.Wrapping = fyne.TextWrapWord

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		grid,
		desc,
	)
}

func (s *sheet) boolToStr(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
