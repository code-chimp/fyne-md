package ui

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestScreen_MakeUI(t *testing.T) {
	testText := "Hello, World!"
	var screen Screen

	edit, preview := screen.MakeUI()

	test.Type(edit, testText)

	if preview.String() != testText {
		t.Errorf("Expected %s, but got %s", testText, preview.String())
	}
}
