package stage

import (
	"github.com/damien-lloyd/gltext"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"kami/render/text"
)

type Welcome struct {

}

func (w Welcome) Name() string {
	return "Welcome Screen"
}

var (
	txts []*gltext.Text
)

func (w Welcome) Load() {
	scaleMin, scaleMax := float32(1.0), float32(1.1)
	for _, str := range []string{"Hello", "there"} {
		txt := gltext.NewText(text.DefaultFont, scaleMin, scaleMax)
		txt.SetString(str)
		txt.SetColor(mgl32.Vec3{1, 1, 1}) //white
		txts = append(txts, txt)
	}
}

func (w Welcome) Draw(window *glfw.Window, delta float32) {
	for index, txt := range txts {
		text.Draw(float32(50), float32(80 + index*50), txt)
	}
}

func (w Welcome) Dispose() {
	for _, txt := range txts {
		txt.Release()
	}
}
