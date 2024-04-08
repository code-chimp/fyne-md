package main

import (
	"fyne-md/internal/ui"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_RunApp(t *testing.T) {
	testTitle := "Markdown Editor"
	testInput := "I'm a little teapot"

	var screen ui.Screen
	screen.AppTitle = testTitle
	testApp := test.NewApp()
	testWin := testApp.NewWindow(screen.AppTitle)

	edit, preview := screen.MakeUI()
	screen.CreateMenuItems(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))

	testApp.Run()

	test.Type(edit, testInput)
	if preview.String() != testInput {
		t.Errorf("Expected %s, but got %s", testInput, preview.String())
	}
}
