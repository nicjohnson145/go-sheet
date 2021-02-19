package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type sheet struct {
	window fyne.Window
}

func newSheet() *sheet {
	return &sheet{}
}

func (s *sheet) loadUI(app fyne.App) {
	s.window = app.NewWindow("Go-Sheet")
	s.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			s.basicStats(),
			s.setupTabs(),
		),
	)
	s.window.Show()
}

func (s *sheet) basicStats() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			s.attributeBlock("Str", 8, "-1"),
			s.attributeBlock("Dex", 10, "+0"),
			s.attributeBlock("Con", 10, "+0"),
			s.attributeBlock("Int", 14, "+2"),
			s.attributeBlock("Wis", 14, "+2"),
			s.attributeBlock("Cha", 18, "+4"),
			s.basicCard("AC", "13"),
			s.basicCard("Proficiency", "+2"),
		),
	)
}

func (s *sheet) attributeBlock(name string, raw int, mod string) fyne.CanvasObject {
	return s.basicCard(name, fmt.Sprintf("%v (%v)", raw, mod))
}

func (s *sheet) basicCard(name string, val string)  *widget.Card {
	return widget.NewCard(
		"",
		"",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				widget.NewLabel(name),
			),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				widget.NewLabel(val),
			),
		),
	)
}

func (s *sheet) setupTabs() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Skills", s.skillTab()),
		container.NewTabItem("Spells", s.spellTab()),
		container.NewTabItem("Weapons", s.weaponsTab()),
		container.NewTabItem("Features", s.featuresTab()),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}

func (s *sheet) skillTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from skill tab")
}

func (s *sheet) spellTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from spell tab")
}

func (s *sheet) weaponsTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from weapons tab")
}

func (s *sheet) featuresTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from features tab")
}
