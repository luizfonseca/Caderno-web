package main

import (
	"caderno/src/interactors"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func hasGit() string {
	return "git status"
}

func main() {
	interactors.CreateGit()

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "Caderno",
		JS:        js,
		CSS:       css,
		Colour:    "#fff",
		Resizable: true,
	})
	app.Bind(hasGit)
	app.Run()
}
