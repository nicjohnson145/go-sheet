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

	s := newSheet(path)
	err := s.loadSheet()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.loadUI(app)
	app.Run()
}
