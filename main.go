package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	s, err := newSheet(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	s.loadUI(app)
	app.Run()
}
