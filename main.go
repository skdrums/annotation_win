package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/skdrums/annotation_win/screens"
)

func main() {
	// implement
	app := app.New()

	w := app.NewWindow("Annotation")
	w.Resize(fyne.NewSize(400, 400))

	ann := screens.NewAnnotation()
	de := ann.LoadDirectoryEdit()
	ia := ann.LoadImageAnnotation(app, w)

	//　KeyMapping

	// output
	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: "DirectoryEdit", Content: de},
		&widget.TabItem{Text: "ImageAnnotation", Content: ia},
	)
	tabs.SetTabLocation(widget.TabLocationTop)
	w.SetContent(tabs)

	w.ShowAndRun()
}
