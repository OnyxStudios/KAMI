package text

import (
	"bytes"
	"fmt"
	"github.com/damien-lloyd/gltext"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/image/math/fixed"
	"kami/render"
	"kami/util"
)

const (
	fontconfigPath  = "fontconfigs"
	defaultFontName = "mono"
)

var (
	DefaultFont *gltext.Font
	fonts       []*gltext.Font
)

func SetWindowSize(width, height float32) {
	for _, font := range fonts {
		font.ResizeWindow(width, height)
		font.OrthographicMatrix = mgl32.Ortho2D(0.0, width, height, 0.0)
		font.OrthographicMatrix = font.OrthographicMatrix.Mul4(mgl32.Scale3D(1.0, -1.0, 1.0))
	}
}

// TODO make text renderer a struct
// TODO store all possible chars as separate text, to not have to preallocate each string
func Draw(x, y float32, text *gltext.Text) {
	text.SetPosition(mgl32.Vec2{x, -y})
	text.Draw()
}

func LoadFonts() {
	DefaultFont = LoadTTF(defaultFontName, fixed.Int26_6(18), fixed.Int26_6(10)) //18pt
}

func ReleaseFonts() {
	for _, font := range fonts {
		font.Release()
	}
}

func LoadTTF(fontName string, fontScale fixed.Int26_6, runesPerRow fixed.Int26_6) *gltext.Font {
	render.InitGL() //load opengl if it isn't loaded yet
	config, err := gltext.LoadTruetypeFontConfig(fontconfigPath, fontName)
	if err != nil { //no font config found or error, regenerate font config
		fontReader := bytes.NewReader(util.ReadAsset(fmt.Sprintf("fonts/%v.ttf", fontName)))
		runeRanges := make(gltext.RuneRanges, 0)
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x0020, High: 0x7F}) //standard ASCII character range
		config, err = gltext.NewTruetypeFontConfig(fontReader, fontScale, runeRanges, runesPerRow, 5)
		util.CheckErr(err)
		util.CheckErr(config.Save(fontconfigPath, fontName))
	}
	font, err := gltext.NewFont(config)
	util.CheckErr(err)
	fonts = append(fonts, font)
	return font
}
