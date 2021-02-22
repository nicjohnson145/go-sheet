package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) skillTab() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		s.skills(),
		widget.NewSeparator(),
		s.saves(),
	)
}

func (s *sheet) skills() fyne.CanvasObject {
	profs := []fyne.CanvasObject{}
	names := []fyne.CanvasObject{}
	mods := []fyne.CanvasObject{}
	for _, skill := range s.character.AllSkills {
		profs = append(profs, profWidget(func() bool { return s.character.isProficientInSkill(skill) }))
		names = append(names, widget.NewLabel(skill.Name))
		mods = append(mods, widget.NewLabel(s.character.modStringForSkill(skill)))
	}

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			widget.NewLabel("Skills"),
		),
		widget.NewSeparator(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				profs...,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				names...,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				mods...,
			),
		),
	)
}

func (s *sheet) saves() fyne.CanvasObject {
	profs := []fyne.CanvasObject{}
	names := []fyne.CanvasObject{}
	mods := []fyne.CanvasObject{}
	for _, save := range s.character.AllSavingThrows {
		profs = append(profs, profWidget(func() bool { return s.character.isProficientInSave(save) }))
		names = append(names, widget.NewLabel(save))
		mods = append(mods, widget.NewLabel(s.character.modStringForSave(save)))
	}

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			widget.NewLabel("Saves"),
		),
		widget.NewSeparator(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				profs...,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				names...,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				mods...,
			),
		),
	)
}

func profWidget(is func() bool) fyne.CanvasObject {
	if is() {
		return widget.NewIcon(theme.RadioButtonCheckedIcon())
	}

	return widget.NewIcon(theme.RadioButtonIcon())
}
