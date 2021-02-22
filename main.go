package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	s, err := newSheet()
	if err != nil {
		fmt.Print(err)
		return
	}
	s.loadUI(app)
	app.Run()
}
