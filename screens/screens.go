package screens

import "fyne.io/fyne/canvas"

var (
	CurrentDir    = "/Users/s07349/go/src/github.com/skdrums/annotation_win/"
	ImageDefault  = "/Users/s07349/go/src/github.com/skdrums/annotation_win/screens/default_image.png"
	LocalImageDir = "original_image/"
	NewDir        = "new_dir/"
	Key           = []string{"A", "S", "D", "F", "G", "J"}
)

type Annotation struct {
	WorkDir       string
	LocalImageDir string
	NewDir        string
	Labels        map[string]string
	Paths         []string

	ViewImageNum int
	ViewImage    *canvas.Image
}

func NewAnnotation() *Annotation {
	image := canvas.NewImageFromFile(ImageDefault)
	image.FillMode = canvas.ImageFillContain
	return &Annotation{
		WorkDir:       CurrentDir,
		LocalImageDir: LocalImageDir,
		NewDir:        NewDir,
		Labels: map[string]string{
			"F": "cat",
			"D": "dog",
		},
		ViewImageNum: 0,
		ViewImage:    image,
	}
}
