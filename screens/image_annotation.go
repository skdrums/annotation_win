package screens

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"fyne.io/fyne/canvas"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

func (a *Annotation) LoadImageAnnotation(app fyne.App, w fyne.Window) fyne.CanvasObject {
	a.Paths = readAllFiles(a.WorkDir + a.LocalImageDir)

	w.Canvas().SetOnTypedKey(a.Control())
	return fyne.NewContainerWithLayout(layout.NewMaxLayout(), a.ViewImage)
}

func readAllFiles(dir string) []string {
	paths := []string{"nil"}

	ob, err := ioutil.ReadDir(dir)
	if err != nil {
		paths = append(paths, ImageDefault)
		return paths
	}
	for _, o := range ob {
		path := filepath.Join(dir, o.Name())
		if !o.IsDir() {
			paths = append(paths, path)
			continue
		}
		paths = append(paths, readAllFiles(path)...)
	}
	for _, path := range paths {
		fmt.Println(filepath.Base(path))
	}
	return paths
}

func (a *Annotation) Control() func(event *fyne.KeyEvent) {
	// FIXME: image表示されていない場合returnする実装
	return func(event *fyne.KeyEvent) {
		if a.ViewImageNum == len(a.Paths) {
			fmt.Println("All images divided")
			return
		}
		if a.ViewImageNum == 0 {
			a.ViewImageNum++
			a.UpdateViewImage(a.ViewImageNum)
			a.ViewImage.FillMode = canvas.ImageFillContain
			canvas.Refresh(a.ViewImage)
			fmt.Println("num: ", a.ViewImageNum, ` before return l56`)
			return
		}
		switch event.Name {
		case fyne.KeyF:
			fmt.Println("num: ", a.ViewImageNum, ` before a.copyResource(a.Labels["F"])`)
			a.copyResource(a.Labels["F"])
		case fyne.KeyD:
			a.copyResource(a.Labels["D"])
			fmt.Println("num: ", a.ViewImageNum, ` before a.copyResource(a.Labels["D"])`)
		case fyne.KeyS:
			a.copyResource(a.Labels["S"])
		case fyne.KeyA:
			a.copyResource(a.Labels["A"])
		case fyne.KeyG:
			a.copyResource(a.Labels["G"])
		case fyne.KeyJ:
			a.copyResource(a.Labels["J"])
		}
		a.ViewImageNum++
		a.UpdateViewImage(a.ViewImageNum)
		a.ViewImage.FillMode = canvas.ImageFillContain
		canvas.Refresh(a.ViewImage)
	}
}

func (a *Annotation) UpdateViewImage(num int) {
	res, err := fyne.LoadResourceFromPath(a.Paths[num])
	if err != nil {
		fmt.Println("error in screens control")
	}
	a.ViewImage.Resource = res
}

func (a *Annotation) copyResource(label string) {
	fmt.Println("path: ", filepath.Base(a.Paths[a.ViewImageNum]), "num: ", a.ViewImageNum)
	src, err := os.Open(a.Paths[a.ViewImageNum])
	if err != nil {
		panic(err)
	}
	defer src.Close()

	newPath := filepath.Join(a.WorkDir, a.NewDir, label, filepath.Base(a.Paths[a.ViewImageNum]))
	dst, err := os.Create(newPath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}
