package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io"
	"strings"
)

var filter = storage.NewExtensionFileFilter([]string{".md", ".markdown", ".MD"})

type Screen struct {
	AppTitle      string
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

func (screen *Screen) MakeUI() (*widget.Entry, *widget.RichText) {
	// create the edit widget
	edit := widget.NewMultiLineEntry()
	screen.EditWidget = edit

	// create the preview widget
	preview := widget.NewRichTextFromMarkdown("")
	screen.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (screen *Screen) openFunction(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				// user cancelled
				return
			}

			defer read.Close()

			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			screen.EditWidget.SetText(string(data))
			screen.CurrentFile = read.URI()
			screen.SaveMenuItem.Disabled = false

			win.SetTitle(screen.AppTitle + " - " + read.URI().Name())

		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (screen *Screen) saveFunction(win fyne.Window) func() {
	return func() {
		if screen.CurrentFile != nil {
			write, err := storage.Writer(screen.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(screen.EditWidget.Text))
			defer write.Close()
		}
	}
}

func (screen *Screen) saveAsFunction(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				// user cancelled
				return
			}

			// use strings.HasSuffix to verify the file extension is either .md or .markdown
			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") && !strings.HasSuffix(strings.ToLower(write.URI().String()), ".markdown") {
				dialog.ShowInformation("Invalid file extension", "Please use a .md or .markdown file extension", win)
				return
			}

			write.Write([]byte(screen.EditWidget.Text))
			screen.CurrentFile = write.URI()

			defer write.Close()

			screen.SaveMenuItem.Disabled = false

			win.SetTitle(screen.AppTitle + " - " + write.URI().Name())

		}, win)

		saveDialog.SetFileName("Untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}

func (screen *Screen) CreateMenuItems(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", screen.openFunction(win))
	saveMenuItem := fyne.NewMenuItem("Save", screen.saveFunction(win))
	screen.SaveMenuItem = saveMenuItem
	screen.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save As...", screen.saveAsFunction(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
