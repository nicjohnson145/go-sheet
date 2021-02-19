package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	
	s := newSheet()
	s.loadUI(app)
	app.Run()
}
