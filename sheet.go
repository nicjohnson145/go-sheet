package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type sheet struct {
	window    fyne.Window
	character *Character
	sheetPath string
	tabs      *container.AppTabs
}

func newSheet(path string) *sheet {
	s := &sheet{
		sheetPath: path,
	}
	return s
}

func (s *sheet) loadSheet() error {
	char, err := newCharacter(s.sheetPath)
	if err != nil {
		return err
	}

	s.character = char
	return nil
}

func (s *sheet) loadUI(app fyne.App) {
	s.window = app.NewWindow("Go-Sheet")
	s.setMainWinContent()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			s.updateCharacterData()
		}
	}()
	s.window.Show()
}

func (s *sheet) updateCharacterData() {
	char, err := newCharacter(s.sheetPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !reflect.DeepEqual(char, s.character) {
		s.character = char
		fmt.Println("Updating character data")
		tabIdx := s.tabs.CurrentTabIndex()
		s.setMainWinContent()
		s.tabs.SelectTabIndex(tabIdx)
	}
}

func (s *sheet) setMainWinContent() {
	s.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			s.basicStats(),
			s.setupTabs(),
			//fyne.NewContainerWithLayout(
			//    layout.NewHBoxLayout(),
			//    layout.NewSpacer(),
			//    widget.NewButton("Refresh", func() {
			//        s.updateCharacterData()
			//    }),
			//),
		),
	)
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
				widget.NewSeparator(),
				widget.NewLabel(fmt.Sprintf(
					"HP: %v / %v",
					s.character.HitPoints.Current,
					s.character.HitPoints.Max,
				)),
				widget.NewSeparator(),
				widget.NewLabel(fmt.Sprintf(
					"Hit Dice (%v): %v / %v",
					s.character.HitDice.Dice,
					s.character.HitDice.Current,
					s.character.HitDice.Max,
				)),
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

func (s *sheet) basicCard(name string, val string) *widget.Card {
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
	s.tabs = container.NewAppTabs(
		container.NewTabItem("Skills", s.skillTab()),
		container.NewTabItem("Spells", s.spellTab()),
		container.NewTabItem("Weapons", s.weaponsTab()),
		container.NewTabItem("Features", s.featuresTab()),
		container.NewTabItem("Resources", s.resourceTab()),
		container.NewTabItem("Equipment", s.equipmentTab()),
		container.NewTabItem("Consumables", s.consumablesTab()),
		container.NewTabItem("Loot", s.lootTab()),
	)
	s.tabs.SetTabLocation(container.TabLocationTop)
	return s.tabs
}
