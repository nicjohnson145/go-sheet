package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type AddRemoveSetCancelConfig struct {
	Label     string
	Current   Setable
	WriteFunc func()
}

type SetableWrap struct {
	Current *int
}

func (s *SetableWrap) SetVal(i int) {
	*s.Current = i
}

func (s *SetableWrap) CurrentVal() int {
	return *s.Current
}

type Setable interface {
	SetVal(int)
	CurrentVal() int
}

func showAddRemoveSetCancelModal(canvas fyne.Canvas, conf AddRemoveSetCancelConfig) {
	entry := widget.NewEntry()
	modal := widget.NewModalPopUp(nil, canvas)

	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel(conf.Label),
		entry,
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				widget.NewButton(
					"Remove",
					func() {
						val, err := strconv.Atoi(entry.Text)
						if err != nil {
							fmt.Println(err)
							return
						}
						conf.Current.SetVal(conf.Current.CurrentVal() - val)
						conf.WriteFunc()
						modal.Hide()
					},
				),
				widget.NewButton(
					"Set",
					func() {
						val, err := strconv.Atoi(entry.Text)
						if err != nil {
							fmt.Println(err)
							return
						}
						conf.Current.SetVal(val)
						conf.WriteFunc()
						modal.Hide()
					},
				),
				widget.NewButton(
					"Add",
					func() {
						val, err := strconv.Atoi(entry.Text)
						if err != nil {
							fmt.Println(err)
							return
						}
						conf.Current.SetVal(conf.Current.CurrentVal() + val)
						conf.WriteFunc()
						modal.Hide()
					},
				),
			),
		),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			widget.NewButton(
				"Cancel",
				func() {
					modal.Hide()
				},
			),
		),
	)
	modal.Content = content
	modal.Show()
}
