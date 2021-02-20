package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (s *sheet) skillTab(c *Character) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		skills(c),
		widget.NewSeparator(),
		saves(c),
	)
}

func skills(c *Character) fyne.CanvasObject {
	profs := []fyne.CanvasObject{}
	names := []fyne.CanvasObject{}
	mods := []fyne.CanvasObject{}
	for _, skill := range c.AllSkills {
		profs = append(profs, profWidget(func() bool { return c.isProficientInSkill(skill) }))
		names = append(names, widget.NewLabel(skill.Name))
		mods = append(mods, widget.NewLabel(c.modStringForSkill(skill)))
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

func saves(c *Character) fyne.CanvasObject {
	profs := []fyne.CanvasObject{}
	names := []fyne.CanvasObject{}
	mods := []fyne.CanvasObject{}
	for _, save := range c.AllSavingThrows {
		profs = append(profs, profWidget(func() bool { return c.isProficientInSave(save) }))
		names = append(names, widget.NewLabel(save))
		mods = append(mods, widget.NewLabel(c.modStringForSave(save)))
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
	var char string
	if is() {
		char = "X"
	} else {
		char = " "
	}
	return widget.NewLabel(char)
}
