package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type sheet struct {
	window fyne.Window
	character *Character
}

func newSheet() (*sheet, error) {
	char, err := newCharacter("./example.yml")
	if err != nil {
		return nil, err
	}
	s := &sheet{
		character: char,
	}
	return s, nil
}

func (s *sheet) loadUI(app fyne.App) {
	s.window = app.NewWindow("Go-Sheet")
	s.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			s.basicStats(),
			s.setupTabs(),
		),
	)
	// s.window.Resize(fyne.NewSize(760, 910))

	// TODO: remove once size is finalized
	//go func() {
	//    for {
	//        fmt.Println(s.window.Canvas().Size())
	//        time.Sleep(time.Second * 2)
	//    }
	//}()

	s.window.Show()
}

func (s *sheet) basicStats() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				widget.NewLabel(s.character.Name),
				widget.NewSeparator(),
				widget.NewLabel(fmt.Sprintf("%v (%v)", s.character.Class, s.character.Level)),
			),
		),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				s.attributeBlock("Str", s.character.Attributes.Strength),
				s.attributeBlock("Dex", s.character.Attributes.Dexterity),
				s.attributeBlock("Con", s.character.Attributes.Constitution),
				s.attributeBlock("Int", s.character.Attributes.Intelligence),
				s.attributeBlock("Wis", s.character.Attributes.Wisdom),
				s.attributeBlock("Cha", s.character.Attributes.Charisma),
				s.basicCard("AC", strconv.Itoa(s.character.ArmorClass)),
				s.basicCard("Proficiency", fmt.Sprintf("+%v", s.character.Proficiency)),
			),
		),
	)
}

func (s *sheet) attributeBlock(name string, raw int) fyne.CanvasObject {
	mod := s.character.modString(s.character.calcMod(raw))
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
		container.NewTabItem("Equipment", s.equipmentTab()),
		container.NewTabItem("Consumables", s.consumablesTab()),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}

func (s *sheet) spellTab() fyne.CanvasObject {
	return widget.NewLabel("Hello from spell tab")
}

