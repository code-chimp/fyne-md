package main

import (
	"fyne-md/internal/theme"
	"fyne-md/internal/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var cfg ui.Screen

func main() {
	title := "Markdown Editor"

	// create a fyne app
	a := app.New()
	a.Settings().SetTheme(&theme.MyTheme{})

	// create a window for the app
	win := a.NewWindow(title)

	// get the user ui
	cfg.AppTitle = title
	edit, preview := cfg.MakeUI()
	cfg.CreateMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show the window and run the app
	win.Resize(fyne.NewSize(800, 600))
	win.CenterOnScreen()
	win.ShowAndRun()
}
