package main

import (
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
	titleRow := []fyne.CanvasObject{
		widget.NewLabel(name),
		layout.NewSpacer(),
	}
	section := s.character.Spells[name]

	if name != Cantrips {
		titleRow = s.addSlotInfo(titleRow, section)
	}

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			titleRow...,
		),
		s.spellAccordion(section.Spells),
	)
}

func (s *sheet) addSlotInfo(row []fyne.CanvasObject, section SpellSecion) []fyne.CanvasObject {
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
	var concentration string
	if spell.Concentration {
		concentration = "yes"
	} else {
		concentration = "no"
	}

	var ritual string
	if spell.Ritual {
		ritual = "yes"
	} else {
		ritual = "no"
	}

	desc := widget.NewLabel(spell.Desc)
	desc.Wrapping = fyne.TextWrapWord

	return fyne.NewContainerWithLayout(
		layout.NewFormLayout(),
		widget.NewLabel("Range"),
		widget.NewLabel(spell.Range),
		widget.NewLabel("Duration"),
		widget.NewLabel(spell.Duration),
		widget.NewLabel("Concentration"),
		widget.NewLabel(concentration),
		widget.NewLabel("Components"),
		widget.NewLabel(strings.Join(spell.Components, ", ")),
		widget.NewLabel("Casting Time"),
		widget.NewLabel(spell.CastingTime),
		widget.NewLabel("Ritual"),
		widget.NewLabel(ritual),
		widget.NewLabel("Description"),
		desc,
	)
}
