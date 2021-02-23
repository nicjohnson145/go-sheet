package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) skillTab() fyne.CanvasObject {
	c := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		s.skills(),
		widget.NewSeparator(),
		s.saves(),
		widget.NewSeparator(),
		s.infoColumn(),
	)
	return c
}

func (s *sheet) skills() fyne.CanvasObject {
	contents := []fyne.CanvasObject{}
	for _, skill := range s.character.AllSkills {
		contents = append(
			contents,
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				profWidget(func() bool { return s.character.isProficientInSkill(skill) }),
				widget.NewLabel(skill.Name),
				widget.NewLabel(s.character.modStringForSkill(skill)),
			),
		)
	}

	return widget.NewCard(
		"Skills",
		"",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			contents...,
		),
	)
}

func (s *sheet) saves() fyne.CanvasObject {
	contents := []fyne.CanvasObject{}
	for _, save := range s.character.AllSavingThrows {
		contents = append(
			contents,
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				profWidget(func() bool { return s.character.isProficientInSave(save) }),
				widget.NewLabel(save),
				widget.NewLabel(s.character.modStringForSave(save)),
			),
		)
	}

	return widget.NewCard(
		"Saves",
		"",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			contents...,
		),
	)
}

func profWidget(is func() bool) fyne.CanvasObject {
	if is() {
		return widget.NewIcon(theme.RadioButtonCheckedIcon())
	}

	return widget.NewIcon(theme.RadioButtonIcon())
}

func (s *sheet) infoColumn() fyne.CanvasObject {
	contents := []fyne.CanvasObject{}

	contents = append(contents, s.listToLabels("Languages", s.character.Languages))
	if len(s.character.Proficiencies.Armor) > 0 {
		contents = append(contents, s.listToLabels("Armor", s.character.Proficiencies.Armor))
	}
	if len(s.character.Proficiencies.Tools) > 0 {
		contents = append(contents, s.listToLabels("Tools", s.character.Proficiencies.Tools))
	}
	if len(s.character.Proficiencies.Weapons) > 0 {
		contents = append(contents, s.listToLabels("Weapons", s.character.Proficiencies.Weapons))
	}
	if len(s.character.Expertise.Skills) > 0 || len(s.character.Expertise.Tools) > 0 {
		expertise := append(s.character.Expertise.Skills, s.character.Expertise.Tools...)
		contents = append(contents, s.listToLabels("Expertise", expertise))
	}

	return container.NewVScroll(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			contents...,
		),
	)
}

func (s *sheet) listToLabels(title string, items []string) fyne.CanvasObject {
	contents := make([]fyne.CanvasObject, 0, len(items))
	for _, item := range items {
		contents = append(contents, widget.NewLabel(item))
	}

	return widget.NewCard(
		title,
		"",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			contents...,
		),
	)
}
