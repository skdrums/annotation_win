package screens

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (a *Annotation) LoadDirectoryEdit() fyne.CanvasObject {
	// ディレクトリ作成form
	workDir := widget.NewEntry()
	workDir.SetPlaceHolder("Input work directory")

	localImageDir := widget.NewEntry()
	localImageDir.SetPlaceHolder("Input local image directory")

	newDir := widget.NewEntry()
	newDir.SetPlaceHolder("Input new image directory")

	newLabels := make(map[string]*widget.Entry)
	for _, k := range Key {
		newLabels[k] = widget.NewEntry()
		newLabels[k].SetPlaceHolder("Input new directory name. This Dir mapped Key " + k)
	}

	// Apply ボタン
	mkDirButton := widget.NewButton("Create", func() {
		if workDir.Text != "" {
			a.WorkDir = workDir.Text
		}
		if localImageDir.Text != "" {
			a.LocalImageDir = localImageDir.Text
		}
		if newDir.Text != "" {
			a.NewDir = newDir.Text
		}
		for key, label := range newLabels {
			if label.Text == "" {
				continue
			}
			a.Labels[key] = label.Text
		}
		a.createDir()
	})

	// Output
	ob := []fyne.CanvasObject{
		widget.NewLabel("Hello Fyne!"),
		workDir,
		localImageDir,
		newDir,
		mkDirButton,
	}
	for _, label := range newLabels {
		ob = append(ob, label)
	}

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), ob...)
}

func (a *Annotation) createDir() {
	for _, label := range a.Labels {
		if err := os.MkdirAll(filepath.Join(a.WorkDir, a.NewDir, label), 0777); err != nil {
			fmt.Println(err)
		}
	}
}
